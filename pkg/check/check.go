package check

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"text/template"

	"github.com/ringods/pulumi-resource/pkg/models"
)

type Runner struct {
	LogWriter io.Writer
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (r Runner) Run(req models.InRequest) ([]models.Version, error) {
	if err := req.Source.Validate(); err != nil {
		return []models.Version{}, err
	}
	// This Run function uses a real http client.
	client := &http.Client{}
	return r.getNewerVersions(req, client)
}

func (r Runner) getNewerVersions(req models.InRequest, client HttpClient) ([]models.Version, error) {
	url, err := r.getPulumiPlatformUpdatesURL(req)
	if err != nil {
		return []models.Version{}, err
	}

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "token "+req.Source.Token)
	response, err := client.Do(request)
	updates := models.Updates{}
	err = json.NewDecoder(response.Body).Decode(&updates)

	return []models.Version{}, nil
}

func (r Runner) getPulumiPlatformUpdatesURL(req models.InRequest) (string, error) {
	tmpl, err := template.New("platformURL").Parse(models.PulumiPlatformTemplateURL)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, req.Source)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func (r Runner) createResponseFromUpdates(updates models.Updates) models.InResponse {
	return models.InResponse{}
}
