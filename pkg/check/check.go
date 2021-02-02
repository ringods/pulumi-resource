package check

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sort"
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
	platformUpdates, err := r.getUpdatesFromPulumiPlatform(req, client)
	if err != nil {
		return []models.Version{}, err
	}
	response := r.createResponseFromUpdates(req, platformUpdates)
	sort.Slice(response, func(i, j int) bool {
		return response[i].Update < response[j].Update
	})

	return response, nil
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

func (r Runner) getUpdatesFromPulumiPlatform(req models.InRequest, client HttpClient) (models.Updates, error) {
	url, err := r.getPulumiPlatformUpdatesURL(req)
	if err != nil {
		return models.Updates{}, err
	}

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "token "+req.Source.Token)
	response, err := client.Do(request)
	updates := models.Updates{}
	err = json.NewDecoder(response.Body).Decode(&updates)

	return updates, nil
}

func (r Runner) createResponseFromUpdates(req models.InRequest, updates models.Updates) []models.Version {
	succesFullUpdates := filter(updates.Updates, func(update models.Update) bool {
		return (update.Info.Result == "succeeded")
	})
	var newerVersions []models.Update
	if req.Version.Update > 0 {
		newerVersions = filter(succesFullUpdates, func(update models.Update) bool {
			return (update.Version >= req.Version.Update)
		})
	} else {
		newerVersions = succesFullUpdates
	}
	response := []models.Version{}
	for _, v := range newerVersions {
		response = append(response, models.Version{Update: v.Version})
	}
	return response
}

func filter(updates []models.Update, f filterFunc) []models.Update {
	var filtered []models.Update
	for _, update := range updates {
		if f(update) {
			filtered = append(filtered, update)
		}
	}
	return filtered
}

type filterFunc func(models.Update) bool
