# Pulumi Resource Type for Concourse

A [Concourse](http://concourse-ci.org/) resource type that allows jobs to modify IaaS resources via [Pulumi](https://www.pulumi.com/). This resource will work against the Pulumi hosted platform. Cloud based state storage backends are not supported by this resource.

*NOTE:* This resource is currently under development and might still contain some bugs. Use at least v0.3.1, in combination with a Pulumi 3.x based project. All previous development versions are unsupported.

## Community

For usage questions, please post a new message on the [Github Discussions](https://github.com/ringods/pulumi-resource/discussions) board for this project.

If you are not sure whether your problem is a usage problem or a bug, please reach out via the discussions first. Also reach out first via the Discussions if you have a suggestion for an improvement.

Only file a new Github issue when you are really sure there is a bug.

## Using the resource type in your pipeline

To use this resource type in your pipeline, you have to register it under the `resource_types` section in your pipeline:

### Example

```yaml
resource_types:
- name: pulumi
  type: registry-image
  source:
    repository: ghcr.io/ringods/pulumi-resource
    tag: v0.3.1
```

Concourse will now know about the resource type called `pulumi` in your list of resources.

## Source Configuration

Every stack you want to manage in your pipeline needs to be registered as a resource in 
your pipeline. Each stack is identified by the following 3 properties, extended with
authentication credentials:

* `organization`: *Required.* The name of the organization you use on the Pulumi platform.
* `project`: *Required.* The name of your Pulumi project.
* `stack`: *Required.* Name of the stack to manage, e.g. `staging`.
* `token`: *Required.* Access token which will be used to login on the Pulumi platform. Use Concourse [Credential Management](https://concourse-ci.org/creds.html) to keep your token safe!

### Example

```yaml
resources:
- name: my-staging-network
  type: pulumi
  source:
    organization: companyname
    project: network
    stack: staging
    token: pul-XXXXXXXXXXXXXXXXX
```

## Behavior

### `check`: Check for new infrastructure deployments

This uses a non-public REST endpoint on the hosted platform, authenticated with the provided
access token, to check for new revisions of the configured stack. The `check` step filters out
failed deployments and will only provide new versions for successful deploys.

### `in`: Get the details of a new infrastructure revision

You can use a Pulumi resource in a `get` step to act as a trigger for downstream builds.
This triggering will happen once the `check` step finds new succesful deploys of the
configured stack.

The details of this new infrastructure revision are not yet provided for downstream processing.

#### Example

```yaml
- name: after-update-of-my-staging-network
  plan:
  - get: my-staging-network
    trigger: true
  - task: do-something-after-deploying-staging-network-stack
    ...
```

### `out`: run pulumi to deploy the latest infrastructure code

Use a pulumi resource in a `put` step if you want to run pulumi in your infrastructure code via Concourse.

Pulumi can use different language runtimes. The amount of possibilities in language runtime, 
language runtime version as well as additional tools on the side makes it sheer impossible
to provide that all within the container image for this resource type. Therefor, this resource
type is implemented in a way that you can pass your custom runtime image via config.

* `runtime`: should point to a `registry-image` resource containing your specific language runtime,
  version and additional tooling.
* `sources`: should point to an input retrieved via `get` or to an `output` from a previous step.
* `config`: this section may contain any valid configuration which you would normally put in your
  `Pulumi.<stack>.yaml` file. Secrets are *not yet* supported.
* `env`: this section may contain environment variables which are needed. E.g., for authentication

Let's show this at work with an example of a NodeJS based Pulumi stack:

#### Example

Pipeline file:

```yaml
resources:
- name: nodejs14
  type: registry-image
  source:
    repository: node

- name: myinfracode
  type: git
  uri: git@github.com:owner/myinfracode.git
    branch: main

- name: my-staging-network
  type: pulumi
  source:
    organization: companyname
    project: network
    stack: staging
    token: pul-XXXXXXXXXXXXXXXXX

jobs:
- name: update-infra
  plan:
  - get: myinfracode
    trigger: true
  - get: nodejs14
  - task: npm-install
    image: nodejs14
    input_mapping: { code: myinfracode }
    file: code/npm-install.yml
  - put: my-staging-network
    params:
      runtime: nodejs14
      sources: code
      config:
        network:setting1: value1
        network:setting2: value2
      env:
        CLOUD_PROVIDER_CREDENTIALS: |
          {
            "private_key_id": "<private key id string>",
            "private_key": "<private key string here>",
            ...
          }
        CLOUD_PROVIDER_REGION: value1
        CLOUD_PROVIDER_ZONE: value2
```

The example task file:

```yaml
---
platform: linux

inputs:
- name: code

outputs:
- name: code

run:
  path: /bin/bash
  args:
    - -c
    - |
      set -euo pipefail
      cd ./code
      npm ci
```

The resources section contains 3 resources:

* `nodejs14`: a container image for NodeJS 14 and support tools like npm or yarn.
  You can configures this fully to your liking.
* `myinfracode`: a git resource pointing to your Pulumi code in a git repository, 
  e.g. tracking new commits on branch `main`
* `my-staging-network`: the pulumi resource pointing to our staging network stack.

We then create a job which does 4 steps:

* retrieves the new revision of the code in the `get: myinfracode` step.
* fetches your NodeJS runtime image in the `get: nodejs14` step. Do not forget to `get` your runtime
  image in the job. If you forget this step, you will not have it available in your `put` step.
* runs `npm install` first on the retrieved code using the `nodejs14` image as the container
* pass your modified sources as an output. If you don't do this and pass your sources using your `get` resource, you still have your clean sources.
* runs Pulumi (via the Automation API) on your code, using the provided runtime image. The stack config
  set via the Concourse pipeline is mixed with any config already provided in the `Pulumi.<stack>.yaml` file residing in the source repository

## Full Example

To wire all the pieces together, here is a full example combining all the previous snippets together:

```yaml
resource_types:
- name: pulumi
  type: registry-image
  source:
    repository: ghcr.io/ringods/pulumi-resource
    tag: v0.2.0

resources:
- name: nodejs14
  type: registry-image
  source:
    repository: node

- name: myinfracode
  type: git
  uri: git@github.com:owner/myinfracode.git
    branch: master

- name: my-staging-network
  type: pulumi
  source:
    organization: companyname
    project: network
    stack: staging
    token: pul-XXXXXXXXXXXXXXXXX

jobs:
- name: update-infra
  plan:
  - get: myinfracode
    trigger: true
  - get: nodejs14
  - task:
    image: nodejs14
    input_mapping: { code: myinfracode }
    file: code/npm-install.yml
  - put: my-staging-network
    params:
      runtime: nodejs14
      sources: code
      config:
        network:setting1: value1
        network:setting2: value2
      env:
        CLOUD_PROVIDER_CREDENTIALS: |
          {
            "private_key_id": "<private key id string>",
            "private_key": "<private key string here>",
            ...
          }
        CLOUD_PROVIDER_REGION: value1
        CLOUD_PROVIDER_ZONE: value2

- name: after-update-infra
  plan:
  - get: my-staging-network
    trigger: true
  - task: do-something-after-rolling-out-network-stack
    ...
```
