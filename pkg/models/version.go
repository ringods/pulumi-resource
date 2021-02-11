package models

// Version is the struct representing the JSON snippet of a version passed between Concourse and the 3 binaries of this resource
type Version struct {
	Update int `json:"update,string"`
}
