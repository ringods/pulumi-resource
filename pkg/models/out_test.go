package models

import (
	"encoding/json"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/config"
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
	assert.Equal(t, OutParams{}, request.Params)
}

func TestDecodeOutRequestWithFlatParamsList(t *testing.T) {
	jsonRequest := []byte("{ \"source\": " + source + ", \"params\": { \"config\": { \"proj:key1\": \"value1\", \"proj:key2\": \"value2\" }}}")
	request := OutRequest{}
	if err := json.Unmarshal(jsonRequest, &request); err != nil {
		assert.Fail(t, "Failed to unmarshal to OutRequest: %s", err)
	}

	expected := config.Map{}
	expected.Set(config.MustMakeKey("proj", "key1"), config.NewValue("value1"), false)
	expected.Set(config.MustMakeKey("proj", "key2"), config.NewValue("value2"), false)

	assert.NotNil(t, request)
	assert.NotNil(t, request.Params)
	assert.NotEmpty(t, request.Params)
	assert.Equal(t, expected, request.Params.Config)
}

func TestDecodeOutRequestWithFlatParamsListMixedTypes(t *testing.T) {
	jsonRequest := []byte("{ \"source\": " + source + ", \"params\": { \"config\": { \"proj:key1\": \"value1\", \"proj:key2\": 2 }}}")
	request := OutRequest{}
	if err := json.Unmarshal(jsonRequest, &request); err != nil {
		assert.Fail(t, "Failed to unmarshal to OutRequest: %s", err)
	}

	expected := config.Map{}
	expected.Set(config.MustMakeKey("proj", "key1"), config.NewValue("value1"), false)
	expected.Set(config.MustMakeKey("proj", "key2"), config.NewObjectValue("2"), false)

	assert.NotNil(t, request)
	assert.NotNil(t, request.Params)
	assert.NotEmpty(t, request.Params)
	assert.Equal(t, expected, request.Params.Config)
}

func TestDecodeOutRequestWithStructuredConfig(t *testing.T) {
	jsonRequest := []byte("{ \"source\": " + source + ", \"params\": { \"config\": { \"proj:data\": {\"active\":true, \"nums\": [ 1, 2, 3 ] } }}}")
	request := OutRequest{}
	if err := json.Unmarshal(jsonRequest, &request); err != nil {
		assert.Fail(t, "Failed to unmarshal to OutRequest: %s", err)
	}

	expected := config.Map{}
	expected.Set(config.MustMakeKey("proj", "data"), config.NewObjectValue("{\"active\":true,\"nums\":[1,2,3]}"), false)

	assert.NotNil(t, request)
	assert.NotNil(t, request.Params)
	assert.NotEmpty(t, request.Params)
	assert.Equal(t, expected, request.Params.Config)
}

func TestGetConfigMapFlatParamsList(t *testing.T) {
	jsonRequest := []byte("{ \"source\": " + source + ", \"params\": { \"config\": { \"proj:key1\": \"value1\", \"proj:key2\": \"value2\" }}}")
	request := OutRequest{}
	if err := json.Unmarshal(jsonRequest, &request); err != nil {
		assert.Fail(t, "Failed to unmarshal to OutRequest: %s", err)
	}

	expected := auto.ConfigMap{"proj:key1": auto.ConfigValue{Value: "value1", Secret: false}, "proj:key2": auto.ConfigValue{Value: "value2"}}

	assert.NotNil(t, request)
	assert.NotNil(t, request.Params)
	assert.NotEmpty(t, request.Params)
	assert.Equal(t, expected, request.GetConfigMap())
}
func TestGetConfigMapFlatParamsListMixedTypes(t *testing.T) {
	jsonRequest := []byte("{ \"source\": " + source + ", \"params\": { \"config\": { \"proj:key1\": \"value1\", \"proj:key2\": 2 }}}")
	request := OutRequest{}
	if err := json.Unmarshal(jsonRequest, &request); err != nil {
		assert.Fail(t, "Failed to unmarshal to OutRequest: %s", err)
	}

	expected := auto.ConfigMap{"proj:key1": auto.ConfigValue{Value: "value1", Secret: false}, "proj:key2": auto.ConfigValue{Value: "2"}}

	assert.NotNil(t, request)
	assert.NotNil(t, request.Params)
	assert.NotEmpty(t, request.Params)
	assert.Equal(t, expected, request.GetConfigMap())
}
func TestGetConfigMapStructuredConfig(t *testing.T) {
	jsonRequest := []byte("{ \"source\": " + source + ", \"params\": { \"config\": { \"proj:data\": {\"active\":true, \"nums\": [ 1, 2, 3 ] } }}}")
	request := OutRequest{}
	if err := json.Unmarshal(jsonRequest, &request); err != nil {
		assert.Fail(t, "Failed to unmarshal to OutRequest: %s", err)
	}

	expected := auto.ConfigMap{"proj:data": auto.ConfigValue{Value: "{\"active\":true,\"nums\":[1,2,3]}", Secret: false}}

	assert.NotNil(t, request)
	assert.NotNil(t, request.Params)
	assert.NotEmpty(t, request.Params)
	assert.Equal(t, expected, request.GetConfigMap())
}
