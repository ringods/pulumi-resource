package models

import (
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
)

// OutRequest is the struct representing the JSON coming in via stdin on `out` binary
type OutRequest struct {
	Source Source     `json:"source"`
	Params config.Map `json:"params"`
}

// OutResponse is the struct representing the JSON going out via stdout on `out` binary
type OutResponse struct {
	Version  Version         `json:"version"`
	Metadata []MetadataEntry `json:"metadata"`
}
