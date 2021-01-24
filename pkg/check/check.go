package check

import (
	"io"

	"github.com/ringods/pulumi-resource/pkg/models"
)

type Runner struct {
	LogWriter io.Writer
}

func (r Runner) Run(req models.InRequest) ([]models.Version, error) {
	return nil, nil
}
