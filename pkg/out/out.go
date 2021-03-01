package out

import (
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/x/auto"
	"github.com/ringods/pulumi-resource/pkg/models"
)

type Runner struct {
	LogWriter            io.Writer
	PulumiSourceLocation string
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (r Runner) Run(req models.OutRequest) (models.OutResponse, error) {
	if err := req.Source.Validate(); err != nil {
		return models.OutResponse{}, err
	}

	// Run the Pulumi Automation API here.
	return r.deployWithPulumi(req)
}

// deployWithPulumi is the activity deploying the infrastructure stack using Pulumi via the Automation API
func (r Runner) deployWithPulumi(req models.OutRequest) (models.OutResponse, error) {
	ctx := context.Background()
	projectName := req.Source.Project
	stackName := auto.FullyQualifiedStackName(req.Source.Organization, projectName, req.Source.Stack)

	// initialize a stack from the checked out sources
	stack, err := auto.UpsertStackLocalSource(ctx, stackName, r.PulumiSourceLocation) // TODO when to run `npm install` here
	if err != nil {
		return models.OutResponse{}, errors.Wrap(err, "Failed to create the stack")
	}
	// Set the Pulumi stack configuration. These values are usually in file `Pulumi.<stack>.yaml`
	// stack.SetConfig(ctx, "key", auto.ConfigValue{Value: "value"})

	update, err := stack.Up(ctx)
	if err != nil {
		// TODO Add the update version here from the UpdateSummary
		return models.OutResponse{}, errors.Wrap(err, "Failed to run `up`")
	}

	return models.OutResponse{
		Version: models.Version{
			Update: update.Summary.Version,
		},
	}, nil
}