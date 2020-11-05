// Code generated by Autopilot. DO NOT EDIT.

package rollback

import (
	parameters "github.com/dds-sysu/autopilot/test/e2e/canary/pkg/parameters"
)

type Inputs struct {
	Deployments     parameters.Deployments
	VirtualServices parameters.VirtualServices
}

// FindDeployment returns <Deployment, true> if the item is found. else parameters.Deployment{}, false
func (i Inputs) FindDeployment(name, namespace string) (parameters.Deployment, bool) {
	for _, item := range i.Deployments.Items {
		if item.Name == name && item.Namespace == namespace {
			return item, true
		}
	}
	return parameters.Deployment{}, false
}

// FindVirtualService returns <VirtualService, true> if the item is found. else parameters.VirtualService{}, false
func (i Inputs) FindVirtualService(name, namespace string) (parameters.VirtualService, bool) {
	for _, item := range i.VirtualServices.Items {
		if item.Name == name && item.Namespace == namespace {
			return item, true
		}
	}
	return parameters.VirtualService{}, false
}

type Outputs struct {
	Deployments     parameters.Deployments
	VirtualServices parameters.VirtualServices
}
