package models

// OutRequest is the struct representing the JSON coming in via stdin on `out` binary
type OutRequest struct {
	Source Source                 `json:"source"`
	Params map[string]interface{} `json:"params"` // absent on initial request
}

// OutResponse is the struct representing the JSON going out via stdout on `out` binary
type OutResponse struct {
	Version  Version         `json:"version"`
	Metadata []MetadataEntry `json:"metadata"`
}
