package models

// Updates contains the information of a series of stack updates
type Updates struct {
	Updates      []Update `json:"updates"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Total        int      `json:"total"`
}

// Update contains the information of a series of stack updates
type Update struct {
	Info          Info   `json:"info"`
	ID            string `json:"updateID"`
	Version       int    `json:"version"`
	LatestVersion int    `json:"latestVersion"`
	// Ignore the rest for now
}

// Info contains specific information about a single Update
type Info struct {
	Kind   string `json:"kind"`
	Result string `json:"result"`
	// Ignore the rest for now
}
