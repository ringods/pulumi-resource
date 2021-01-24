package config

// Source contains all the rescource configuration attributes
type Source struct {
	Organization string `json:"organization"`
	Project      bool   `json:"project"`
	Stack        bool   `json:"stack"`
	Token        bool   `json:"token"`
}
