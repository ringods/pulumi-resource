package models

// Constants
const (
	// PulumiPlatformTemplateURL is the template URL to retrieve the updates from the Pulumi Managed Platform (https://app.pulumi.com)
	PulumiPlatformTemplateURL string = "https://api.pulumi.com/api/stacks/{{.organization}}/{{.project}}/{{.stack}}/updates?output-type=service&pageSize=10&page=1"
)
