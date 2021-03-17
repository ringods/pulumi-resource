package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	source string = "{ \"organization\": \"ringods\", \"project\": \"mypulumistack\", \"stack\": \"production\", \"token\": \"pul-XXXXXXXXXXXXXXXXX\" }"
)

func TestDecodeOutRequestWithoutParams(t *testing.T) {
	jsonRequest := []byte("{ \"source\": " + source + ", \"params\": {}}")
	request := OutRequest{}
	if err := json.Unmarshal(jsonRequest, &request); err != nil {
		assert.Fail(t, "Failed to unmarshal to OutRequest: %s", err)
	}
	assert.NotNil(t, request)
	assert.NotNil(t, request.Params)
	assert.Empty(t, request.Params)
	assert.Equal(t, map[string]interface{}{}, request.Params)
}

func TestDecodeOutRequestWithFlatParamsList(t *testing.T) {
	jsonRequest := []byte("{ \"source\": " + source + ", \"params\": { \"key1\": \"value1\", \"key2\": \"value2\" }}")
	request := OutRequest{}
	if err := json.Unmarshal(jsonRequest, &request); err != nil {
		assert.Fail(t, "Failed to unmarshal to OutRequest: %s", err)
	}
	assert.NotNil(t, request)
	assert.NotNil(t, request.Params)
	assert.NotEmpty(t, request.Params)
	assert.Equal(t, map[string]interface{}{"key1": "value1", "key2": "value2"}, request.Params)
}

func TestDecodeOutRequestWithFlatParamsListMixedTypes(t *testing.T) {
	jsonRequest := []byte("{ \"source\": " + source + ", \"params\": { \"key1\": \"value1\", \"key2\": 2 }}")
	request := OutRequest{}
	if err := json.Unmarshal(jsonRequest, &request); err != nil {
		assert.Fail(t, "Failed to unmarshal to OutRequest: %s", err)
	}
	assert.NotNil(t, request)
	assert.NotNil(t, request.Params)
	assert.NotEmpty(t, request.Params)
	assert.Equal(t, map[string]interface{}{"key1": "value1", "key2": 2}, request.Params)
}
