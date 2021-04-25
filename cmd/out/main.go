package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ringods/pulumi-resource/pkg/encoder"
	"github.com/ringods/pulumi-resource/pkg/models"
	"github.com/ringods/pulumi-resource/pkg/out"
)

func main() {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stderr, "Current working directory: %s", currentWorkingDirectory)

	req := models.OutRequest{}
	if err := json.NewDecoder(os.Stdin).Decode(&req); err != nil {
		log.Fatalf("Failed to read OutRequest: %s", err)
	}

	cmd := out.Runner{
		LogWriter:            os.Stderr,
		ConcourseBuildFolder: os.Args[1],
	}
	resp, err := cmd.Run(req)
	if err != nil {
		log.Fatal(err)
	}

	if err := encoder.NewJSONEncoder(os.Stdout).Encode(resp); err != nil {
		log.Fatalf("Failed to write OutResponse to stdout: %s", err)
	}

}
