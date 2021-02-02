package check

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/ringods/pulumi-resource/pkg/models"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	// just in case you want default correct return value
	return &http.Response{}, nil
}

func TestActualPulumiPlatformURL(t *testing.T) {
	cmd := Runner{
		LogWriter: os.Stderr,
	}
	req := models.InRequest{
		Source: models.Source{
			Organization: "ringods",
			Project:      "mypulumiproject",
			Stack:        "production",
			Token:        "pul-XXXXXXXXXXXXXXX",
		},
	}

	url, err := cmd.getPulumiPlatformUpdatesURL(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https://api.pulumi.com/api/stacks/ringods/mypulumiproject/production/updates?output-type=service&pageSize=10&page=1", url)
}

func TestGetNewerVersionsForNewStack(t *testing.T) {
	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// do whatever you want
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader("{\"updates\":[]}")),
			}, nil
		},
	}

	cmd := Runner{
		LogWriter: os.Stderr,
	}
	req := models.InRequest{
		Source: models.Source{
			Organization: "ringods",
			Project:      "mypulumiproject",
			Stack:        "production",
			Token:        "pul-XXXXXXXXXXXXXXX",
		},
		Version: models.Version{
			Update: 0,
		},
	}
	versions, err := cmd.getNewerVersions(req, client)
	// no error
	assert.Equal(t, nil, err)
	// No versions to return
	assert.Equal(t, 0, len(versions))
	assert.Equal(t, []models.Version{}, versions)
}

func TestGetNewerVersionsForNewStackWithUpdates(t *testing.T) {
	jsonFile, err := os.Open("../models/updates.json")
	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// do whatever you want
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(jsonFile),
			}, nil
		},
	}

	cmd := Runner{
		LogWriter: os.Stderr,
	}
	req := models.InRequest{
		Source: models.Source{
			Organization: "ringods",
			Project:      "mypulumiproject",
			Stack:        "production",
			Token:        "pul-XXXXXXXXXXXXXXX",
		},
		Version: models.Version{
			Update: 0,
		},
	}
	versions, err := cmd.getNewerVersions(req, client)
	// no error
	assert.Equal(t, nil, err)
	// No versions to return
	assert.Equal(t, 3, len(versions))
	assert.Equal(t, []models.Version{{Update: 76}, {Update: 77}, {Update: 79}}, versions)
}

func TestGetNewerVersionsForExistingStackWithUpdates(t *testing.T) {
	jsonFile, err := os.Open("../models/updates.json")
	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// do whatever you want
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(jsonFile),
			}, nil
		},
	}

	cmd := Runner{
		LogWriter: os.Stderr,
	}
	req := models.InRequest{
		Source: models.Source{
			Organization: "ringods",
			Project:      "mypulumiproject",
			Stack:        "production",
			Token:        "pul-XXXXXXXXXXXXXXX",
		},
		Version: models.Version{
			Update: 77,
		},
	}
	versions, err := cmd.getNewerVersions(req, client)
	// no error
	assert.Equal(t, nil, err)
	// No versions to return
	assert.Equal(t, 2, len(versions))
	assert.Equal(t, []models.Version{{Update: 77}, {Update: 79}}, versions)
}

func TestGetNewerVersionsForExistingStackWithoutUpdates(t *testing.T) {
	jsonFile, err := os.Open("../models/updates.json")
	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// do whatever you want
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(jsonFile),
			}, nil
		},
	}

	cmd := Runner{
		LogWriter: os.Stderr,
	}
	req := models.InRequest{
		Source: models.Source{
			Organization: "ringods",
			Project:      "mypulumiproject",
			Stack:        "production",
			Token:        "pul-XXXXXXXXXXXXXXX",
		},
		Version: models.Version{
			Update: 79,
		},
	}
	versions, err := cmd.getNewerVersions(req, client)
	// no error
	assert.Equal(t, nil, err)
	// No versions to return
	assert.Equal(t, 1, len(versions))
	assert.Equal(t, []models.Version{{Update: 79}}, versions)
}

func readUpdatesFromFile(t *testing.T) models.Updates {
	jsonFile, err := os.Open("../models/updates.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		assert.Fail(t, err.Error())
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	updateContent, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Updates array
	var updates models.Updates

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(updateContent, &updates)

	return updates
}

func TestCreateInResponseFromUpdates(t *testing.T) {
	req := models.InRequest{
		Source: models.Source{
			Organization: "ringods",
			Project:      "mypulumiproject",
			Stack:        "production",
			Token:        "pul-XXXXXXXXXXXXXXX",
		},
	}

	updates := readUpdatesFromFile(t)

	cmd := Runner{
		LogWriter: os.Stderr,
	}
	response := cmd.createResponseFromUpdates(req, updates)
	assert.NotNil(t, response)
}
