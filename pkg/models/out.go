package models

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/config"
)

type OutParams struct {
	Sources string     `json:"sources"`
	Runtime string     `json:"runtime"`
	Config  config.Map `json:"config"`
}

// OutRequest is the struct representing the JSON coming in via stdin on `out` binary
type OutRequest struct {
	Source Source    `json:"source"`
	Params OutParams `json:"params"`
}

// OutResponse is the struct representing the JSON going out via stdout on `out` binary
type OutResponse struct {
	Version  Version         `json:"version"`
	Metadata []MetadataEntry `json:"metadata"`
}

// GetConfigMap transforms the core Pulumi config.Map to the ConfigMap type from the Automation API
func (req *OutRequest) GetConfigMap() auto.ConfigMap {
	cfgMap := auto.ConfigMap{}

	cfg := req.Params.Config
	// Iterate over all keys
	for key, val := range cfg {
		v, _ := val.Value(nil)
		cfgMap[key.String()] = auto.ConfigValue{Value: v}
	}
	return cfgMap
}

// ExtendPathWithRuntime extends the PATH env value by adding paths to the runtime binaries
func (req *OutRequest) ExtendPathWithRuntime(buildDirectory string, currentPathEnv string) string {
	// Prepend these paths:
	// ${buildDirectory}/${runtime}/rootfs/bin
	// ${buildDirectory}/${runtime}/rootfs/usr/bin
	// ${buildDirectory}/${runtime}/rootfs/usr/local/bin
	baseExtension := fmt.Sprintf("%s/%s/rootfs", buildDirectory, req.Params.Runtime)
	extension := fmt.Sprintf("%s/bin:%s/usr/bin:%s/usr/local/bin", baseExtension, baseExtension, baseExtension)

	return fmt.Sprintf("%s:%s", extension, currentPathEnv)
}

// GetSourceLocation calculates the Concourse specific path the to Pulumi sources.
func (req *OutRequest) GetSourceLocation(buildDirectory string) string {
	return fmt.Sprintf("%s/%s", buildDirectory, req.Params.Sources)
}
