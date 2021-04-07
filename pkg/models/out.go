package models

import (
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v2/go/x/auto"
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
