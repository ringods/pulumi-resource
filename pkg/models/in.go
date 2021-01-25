package models

// InRequest is the struct representing the JSON coming in via stdin on `check` and `in` binaries
type InRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version,omitempty"` // absent on initial request
}

// InResponse is the struct representing the JSON going out via stdout on `check` and `in` binaries
type InResponse struct {
	Version  Version  `json:"version"`
	Metadata Metadata `json:"metadata"`
}
