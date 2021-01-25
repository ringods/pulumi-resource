package models

import (
	"errors"
)

// Source contains all the rescource configuration attributes
type Source struct {
	Organization string `json:"organization"`
	Project      string `json:"project"`
	Stack        string `json:"stack"`
	Token        string `json:"token"`
}

// Validate all the source configuration parameters
func (s Source) Validate() error {
	if s.Organization == "" {
		return errors.New("parameter `organization` can not be the empty string")
	}

	if s.Project == "" {
		return errors.New("parameter `project` can not be the empty string")
	}

	if s.Stack == "" {
		return errors.New("parameter `stack` can not be the empty string")
	}

	if s.Token == "" {
		return errors.New("the access `token` can not be the empty string")
	}

	return nil
}
