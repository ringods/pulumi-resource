package check

import (
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

func TestCheckValidSource(t *testing.T) {
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
	}
	versions, err := cmd.getNewerVersions(req, client)
	assert.Equal(t, nil, err)
	assert.Equal(t, []models.Version{}, versions)
}
