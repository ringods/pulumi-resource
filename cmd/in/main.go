package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ringods/pulumi-resource/pkg/encoder"
	"github.com/ringods/pulumi-resource/pkg/in"
	"github.com/ringods/pulumi-resource/pkg/models"
)

func main() {
	req := models.InRequest{}
	if err := json.NewDecoder(os.Stdin).Decode(&req); err != nil {
		log.Fatalf("Failed to read InRequest: %s", err)
	}

	cmd := in.Runner{
		LogWriter: os.Stderr,
	}
	resp, err := cmd.Run(req)
	if err != nil {
		log.Fatal(err)
	}

	if err := encoder.NewJSONEncoder(os.Stdout).Encode(resp); err != nil {
		log.Fatalf("Failed to write Versions to stdout: %s", err)
	}

}
