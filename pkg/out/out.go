package out

import (
	"io"
	"net/http"

	"github.com/ringods/pulumi-resource/pkg/models"
)

type Runner struct {
	LogWriter io.Writer
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (r Runner) Run(req models.OutRequest) (models.OutResponse, error) {
	if err := req.Source.Validate(); err != nil {
		return models.OutResponse{}, err
	}

	// Run the Pulumi Automation API here.

	return models.OutResponse{
		Version:  models.Version{Update: 1}, // TODO Must be implemented from what the Pulumi Automation API returns
		Metadata: []models.MetadataEntry{},
	}, nil
}
