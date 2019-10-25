package meshrouter

import (
	"context"
	"fmt"
	"github.com/solo-io/autopilot/examples/canary/lib/utils"
	v1 "github.com/solo-io/autopilot/examples/canary/pkg/apis/canaries/v1"
	glookubev1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"strings"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	solokitcore "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	solokiterror "github.com/solo-io/solo-kit/pkg/errors"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GlooRouter is managing Istio virtual services
type GlooRouter struct {
	ezKube              utils.EzKube
	logger              *zap.SugaredLogger
	upstreamDiscoveryNs string
}

func NewGlooRouter(provider string, logger *zap.SugaredLogger, ezKube utils.EzKube) (*GlooRouter, error) {

	upstreamDiscoveryNs := ""
	if strings.HasPrefix(provider, "gloo:") {
		upstreamDiscoveryNs = strings.TrimPrefix(provider, "gloo:")
	}

	return NewGlooRouterWithClient(ezKube, upstreamDiscoveryNs, logger), nil
}

func NewGlooRouterWithClient(ezKube utils.EzKube, upstreamDiscoveryNs string, logger *zap.SugaredLogger) *GlooRouter {
	if upstreamDiscoveryNs == "" {
		upstreamDiscoveryNs = "gloo-system"
	}
	return &GlooRouter{ezKube: ezKube, logger: logger, upstreamDiscoveryNs: upstreamDiscoveryNs}
}

// Reconcile creates or updates the Istio virtual service
func (gr *GlooRouter) Reconcile(ctx context.Context, canary *v1.Canary) error {
	// do we have routes already?
	if _, _, _, err := gr.GetRoutes(ctx, canary); err == nil {
		// we have routes, no need to do anything else
		return nil
	} else if solokiterror.IsNotExist(err) {
		return gr.SetRoutes(ctx, canary, 100, 0, false)
	} else {
		return err
	}
}

// GetRoutes returns the destinations weight for primary and canary
func (gr *GlooRouter) GetRoutes(ctx context.Context, canary *v1.Canary) (
	primaryWeight int,
	canaryWeight int,
	mirrored bool,
	err error,
) {
	targetName := canary.Spec.TargetRef.Name
	ug := &glookubev1.UpstreamGroup{ObjectMeta: metav1.ObjectMeta{
		Name:      targetName,
		Namespace: canary.Namespace,
	}}
	if err = gr.ezKube.Get(ctx, ug); err != nil {
		return
	}

	dests := ug.Spec.GetDestinations()
	for _, dest := range dests {
		if dest.GetDestination().GetUpstream().Name == upstreamName(canary.Namespace, fmt.Sprintf("%s-primary", targetName), canary.Spec.Service.Port) {
			primaryWeight = int(dest.Weight)
		}
		if dest.GetDestination().GetUpstream().Name == upstreamName(canary.Namespace, fmt.Sprintf("%s-canary", targetName), canary.Spec.Service.Port) {
			canaryWeight = int(dest.Weight)
		}
	}

	if primaryWeight == 0 && canaryWeight == 0 {
		err = fmt.Errorf("RoutingRule %s.%s does not contain routes for %s-primary and %s-canary",
			targetName, canary.Namespace, targetName, targetName)
	}

	mirrored = false

	return
}

// SetRoutes updates the destinations weight for primary and canary
func (gr *GlooRouter) SetRoutes(
	ctx context.Context,
	canary *v1.Canary,
	primaryWeight int,
	canaryWeight int,
	mirrored bool,
) error {
	targetName := canary.Spec.TargetRef.Name

	if primaryWeight == 0 && canaryWeight == 0 {
		return fmt.Errorf("RoutingRule %s.%s update failed: no valid weights", targetName, canary.Namespace)
	}

	destinations := []*gloov1.WeightedDestination{}
	destinations = append(destinations, &gloov1.WeightedDestination{
		Destination: &gloov1.Destination{
			DestinationType: &gloov1.Destination_Upstream{
				Upstream: &solokitcore.ResourceRef{
					Name:      upstreamName(canary.Namespace, fmt.Sprintf("%s-primary", targetName), canary.Spec.Service.Port),
					Namespace: gr.upstreamDiscoveryNs,
				},
			},
		},
		Weight: uint32(primaryWeight),
	})

	destinations = append(destinations, &gloov1.WeightedDestination{
		Destination: &gloov1.Destination{
			DestinationType: &gloov1.Destination_Upstream{
				Upstream: &solokitcore.ResourceRef{
					Name:      upstreamName(canary.Namespace, fmt.Sprintf("%s-canary", targetName), canary.Spec.Service.Port),
					Namespace: gr.upstreamDiscoveryNs,
				},
			},
		},
		Weight: uint32(canaryWeight),
	})

	upstreamGroup := &glookubev1.UpstreamGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      canary.Spec.TargetRef.Name,
			Namespace: canary.Namespace,
		},
		Spec: gloov1.UpstreamGroup{
			Destinations: destinations,
		},
	}

	return gr.writeUpstreamGroupRuleForCanary(ctx, canary, upstreamGroup)
}

func (gr *GlooRouter) writeUpstreamGroupRuleForCanary(ctx context.Context, canary *v1.Canary, ug *glookubev1.UpstreamGroup) error {
	targetName := canary.Spec.TargetRef.Name

	oldUg := &glookubev1.UpstreamGroup{ObjectMeta: metav1.ObjectMeta{
		Name:      ug.Namespace,
		Namespace: ug.Name,
	}}
	if err := gr.ezKube.Get(ctx, oldUg); err != nil {
		if errors.IsNotFound(err) {
			gr.logger.With("canary", fmt.Sprintf("%s.%s", canary.Name, canary.Namespace)).
				Infof("UpstreamGroup %s created", ug.Name)
		} else {
			return fmt.Errorf("UpstreamGroup %s.%s read failed: %v", targetName, canary.Namespace, err)
		}
		return err
	} else {
		ug.ResourceVersion = oldUg.ResourceVersion
		// if the old and the new one are equal, no need to do anything.
		oldUg.Status = solokitcore.Status{}
		if oldUg.Spec.Equal(ug.Spec) {
			return nil
		}
	}

	return gr.ezKube.Ensure(ctx, ug)
}

func upstreamName(serviceNamespace, serviceName string, port int32) string {
	return fmt.Sprintf("%s-%s-%d", serviceNamespace, serviceName, port)
}