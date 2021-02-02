package in

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

func (r Runner) Run(req models.InRequest) (models.InResponse, error) {
	if err := req.Source.Validate(); err != nil {
		return models.InResponse{}, err
	}

	return models.InResponse{
		Version:  req.Version,
		Metadata: []models.MetadataEntry{},
	}, nil
}
