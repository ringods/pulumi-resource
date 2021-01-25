package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidSource(t *testing.T) {
	validSource := Source{
		Organization: "ringods",
		Project:      "mypulumistack",
		Stack:        "production",
		Token:        "pul-XXXXXXXXXXXXXXXXX",
	}
	assert.Equal(t, nil, validSource.Validate())
}

func TestCheckSourceMissingOrganization(t *testing.T) {
	validSource := Source{
		Organization: "",
		Project:      "mypulumistack",
		Stack:        "production",
		Token:        "pul-XXXXXXXXXXXXXXXXX",
	}
	assert.Equal(t, "parameter `organization` can not be the empty string", validSource.Validate().Error())
}

func TestCheckSourceMissingProject(t *testing.T) {
	validSource := Source{
		Organization: "ringods",
		Project:      "",
		Stack:        "production",
		Token:        "pul-XXXXXXXXXXXXXXXXX",
	}
	assert.Equal(t, "parameter `project` can not be the empty string", validSource.Validate().Error())
}

func TestCheckSourceMissingStack(t *testing.T) {
	validSource := Source{
		Organization: "ringods",
		Project:      "mypulumistack",
		Stack:        "",
		Token:        "pul-XXXXXXXXXXXXXXXXX",
	}
	assert.Equal(t, "parameter `stack` can not be the empty string", validSource.Validate().Error())
}
func TestCheckSourceMissingToken(t *testing.T) {
	validSource := Source{
		Organization: "ringods",
		Project:      "mypulumistack",
		Stack:        "production",
		Token:        "",
	}
	assert.Equal(t, "the access `token` can not be the empty string", validSource.Validate().Error())
}
