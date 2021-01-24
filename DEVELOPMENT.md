# Development

The code is structured similarly to the [`terraform-resource`](https://github.com/ljfranklin/terraform-resource)

## Building the software

Building the software is easy. Using the provided `GNUmakefile` run

```sh
$ make

Usage:
  make 

Dependencies
  deps-go          Install Go dependencies
  deps             Install all dependencies

Build
  build-in         Build the `in` binary
  build-out        Build the `out` binary
  build-check      Build the `check` binary
  build            Build all binaries

Testing
  test             Run all tests

Helpers
  help             Display this help
```

To see information about the available targets. The main targets are `build` and `test`. These are also triggered in the Github pipelines.

## Add / Updating dependencies

This project uses Go modules. See [here](https://golang.cafe/blog/upgrade-dependencies-golang.html) for information on bumping dependencies

## Testing your changes

Together with your changes, make sure you add the proper unit tests. Then run `make test` to verify the code changes.
Once ready, create and submit a Pull Request to submit your changes.

## Design of the software

First of all, make sure you understand the function of the 3 scripts or binaries `check`, `in` and `out`. The documentation on this can be found on the [Concourse website](https://concourse-ci.org/implementing-resource-types.html).

### check

The `check` binary for this resource type receives a JSON snippet via `stdin`:

```json
{
  "source": {
    "organization": "ringods",
    "project": "mypulumicode",
    "stack": "production",
    "token": "pul-XXXXXXXXXXXXXXXXXXX"
  },
  "version": { "update": "44" }
}
```

The `source` section is passed verbatim from the resource configuration in your Concourse pipeline.

The `version` section returns the identifier of the last update for this stack.

Using the provided access `token`, the code will retrieve the results from this URL:

```
https://api.pulumi.com/api/stacks/{organization}/{project}/{stack}/updates?output-type=service&pageSize=10&page=1
```

The check binary will process the results and calculate the list of versions that are newer than the current one. The result could be like this:

```json
[
  { "update": "44" },
  { "update": "46" },
  { "update": "47" }
]
```

This example shows that there are newer infrastructure updates. We may only trigger downstream builds when upstream runs were succesful. As such, the `check` binary will filter out the failed updates from the range. In this example, update 45 failed and is not returned.


### in

TODO

### out

TODO