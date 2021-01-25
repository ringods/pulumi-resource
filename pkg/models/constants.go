package models

// Constants
const (
	// PulumiPlatformTemplateURL is the template URL to retrieve the updates from the Pulumi Managed Platform (https://app.pulumi.com)
	// Placeholders should match the members of models.Source
	PulumiPlatformTemplateURL string = "https://api.pulumi.com/api/stacks/{{.Organization}}/{{.Project}}/{{.Stack}}/updates?output-type=service&pageSize=10&page=1"
)
