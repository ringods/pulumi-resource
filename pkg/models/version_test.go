package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalVersion(t *testing.T) {
	expectedVersion := Version{Update: 1}
	jsonVersion := `{"update":"1"}`
	version := Version{}
	json.Unmarshal([]byte(jsonVersion), &version)
	assert.Equal(t, expectedVersion, version)
}

func TestMarshalVersion(t *testing.T) {
	expectedJSONVersion := []byte(`{"update":"1"}`)
	version := Version{Update: 1}

	actualJSONVersion, _ := json.Marshal(version)
	assert.Equal(t, expectedJSONVersion, actualJSONVersion)
}
