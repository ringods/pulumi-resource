package models

// InRequest is the struct representing the JSON coming in via stdin on `check` and `in` binaries
type InRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version,omitempty"` // absent on initial request
}

type MetadataEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type InResponse struct {
	Version  Version         `json:"version"`
	Metadata []MetadataEntry `json:"metadata"`
}
