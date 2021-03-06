---
title: "Autopilot SDK Structure"
description: "Overview of the design and structure of the Autopilot SDK"
weight: 2
---

The source code for the Autopilot SDK is composed of 3 core directories:

- `cli`: contains source code for the `ap` CLI. 
- `codegen`: contains the source code for generating code, Kubernetes manifests, and other project files necessary to build and deploy the operator. `codegen` is invoked by the `ap generate` command, but can be invoked manually (see [`debuggable_generate.go`](https://github.com/dds-sysu/autopilot/blob/master/test/e2e/debuggable_generate.go) for an example).
- `pkg`: libraries used for running Autopilot Operators.

### CLI Design

The `ap` CLI is designed to manage the full lifecycle of the Operator, from code generation, build, and deployment. 

Autopilot Operators can be built and deployed via standard means (e.g. `go build`, `docker`, `kubectl apply`). The `ap` CLI is designed to simplify this process, without being a required component.

### Codegen Design

The `codegen` package contains the root `Run`, `Load` and `Generate` functions which are used to generate code and other project files from the `autopilot.yaml` file.

It contains subpackages `model`, `templates`, and `util`.

- `model` - contains the internal model of project data parsed from the `autopilot.yaml` as well as utility functions which are used to render templates for generated files.

- `templates` - contains the templates used to render generated code, deployment, and project files. Templates are organized by type, with `code` containing all templates for generated `.go` files, and `deploy` containing templates for generated Kubernetes manifests.

For the full list of files generated by Autopilot, see [`generate.go#projectFiles`](https://github.com/dds-sysu/autopilot/blob/master/codegen/generate.go#L145) and [`generate.go#phaseFiles`](https://github.com/dds-sysu/autopilot/blob/master/codegen/generate.go#L213).

### Pkg Design

`pkg` contains the core libraries for running Autopilot Operators. It is broken up into the following subpackages:

- `config`: defaults and helper functions for loading the Autopilot Operator config. Read more about the Autopilot Operator config [in the reference documentation]({{< versioned_link_path fromRoot="/reference/api/api_v1">}}#autopilot-operator.proto).
- `defaults`: defaults core to the system. Default file location of `autopilot.yaml` as well as other variables (which can be overridden in `init()` functions).
- `ezkube`: `ezkube` contains a client which is a convenience wrapper for the dynamic `client.Client` of the [controller-runtime library](https://github.com/kubernetes-sigs/controller-runtime/blob/master/pkg/client/interfaces.go#L104). It adds convenience functions for operators such as the `Ensure` function which applies resources to Kubernetes, setting owner references and retrying on resource conflicts.
- `metrics`: defines the base `metrics.Client` on which generated metrics code is based. The implemented metrics client is designed primarily for querying Prometheus.
- `run`: contains the main entrypoint for Autopilot Operators. The Operator's generated `main.go` calls the `run.Run` function which runs the user's scheduler [`Scheduler`](https://github.com/dds-sysu/autopilot/blob/master/codegen/templates/scheduler.gotmpl).
- `scheduler`: contains utilities and shared code for the generated `scheduler.go`, which is responsible for calling the user-defined *workers*.
- `utils`: utilities used in various places in generated code.
- `version`: small package containing the current version of Autopilot itself. This version is stored as a variable which is set at compile-time by Autopilot's `Makefile`.
