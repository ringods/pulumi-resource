# Pulumi Resource Type for Concourse

A [Concourse](http://concourse-ci.org/) resource type that allows jobs to modify IaaS resources via [Pulumi](https://www.pulumi.com/). This resource will work against the Pulumi hosted platform. Cloud based state storage backends are not supported by this resource.

*NOTE:* This resource is currently under development and isn't yet fully functional. As of v0.0.5, it should be usable to function as an resource you can `get`. `put` is under development.

## Community

For usage questions, please post a new message on the [Github Discussions](https://github.com/ringods/pulumi-resource/discussions) board for this project.

If you are not sure whether your problem is a usage problem or a bug, please reach out via the discussions first. Also reach out first via the Discussions if you have a suggestion for an improvement.

Only file a new Github issue when you are really sure there is a bug.

## Source Configuration

* `organization`: *Required.* The name of the organization you use on the Pulumi platform.

* `project`: *Required.* The name of your Pulumi project.

* `stack`: *Required.* Name of the stack to manage, e.g. `staging`.

* `token`: *Required.* Access token which will be used to login on the Pulumi platform.

#### Source Example

```yaml
resource_types:
- name: pulumi
  type: docker-image
  source:
    repository: ghcr.io/ringods/pulumi-resource
    tag: v0.0.5

resources:
  - name: myinfracode
    type: pulumi
    source:
      organization: companyname
      project: network
      stack: staging
      token: pul-XXXXXXXXXXXXXXXXX

jobs:
  - name: nextstack
    plan:
      - get: myinfracode
        trigger: true
      - task: do-something-after-rolling-out-network-stack
        ...
```
