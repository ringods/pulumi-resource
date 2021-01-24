## This is a self-documented Makefile. For usage information, run `make help`:
## For more information, refer to https://www.thapaliya.com/en/writings/well-documented-makefiles/
##
## Structure for a Go project with multiple binaries: https://stackoverflow.com/a/37681238

.DEFAULT_GOAL:=help
SHELL:=/bin/bash

GO = GO111MODULE=on go
GO_FILES ?= ./pkg/...

all: deps build

##@ Dependencies

deps-go: ## Install Go dependencies

deps: deps-go ## Install all dependencies
	$(GO) install -v ./cmd/in ./cmd/out ./cmd/check

##@ Build

build-in: deps ## Build the `in` binary
	$(GO) build -o bin/in ./cmd/in

build-out: deps ## Build the `out` binary
	$(GO) build -o bin/out ./cmd/out

build-check: deps ## Build the `check` binary
	$(GO) build -o bin/check ./cmd/check

build: build-in build-out build-check ## Build all binaries

##@ Testing

test: ## Run all tests
	$(GO) test -v $(GO_FILES)

##@ Helpers

.PHONY: help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
