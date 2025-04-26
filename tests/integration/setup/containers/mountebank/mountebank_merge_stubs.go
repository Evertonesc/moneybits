package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Imposter struct {
	Port     int               `json:"port"`
	Protocol string            `json:"protocol"`
	Stubs    []json.RawMessage `json:"stubs"`
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	stubsDir := filepath.Join(dir, "stubs")
	outputFile := filepath.Join(dir, "imposter.json")

	imposter := Imposter{
		Port:     8090,
		Protocol: "http",
		Stubs:    make([]json.RawMessage, 0),
	}

	stubs, err := walkDirs(stubsDir)
	if err != nil {
		fmt.Printf("Error processing stubs directory: %v\n", err)
		os.Exit(1)
	}

	imposter.Stubs = stubs

	output, err := json.MarshalIndent(imposter, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling imposter: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(outputFile, output, 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Merged imposter file created at: %s\n", outputFile)
}

func walkDirs(dir string) ([]json.RawMessage, error) {
	var stubs []json.RawMessage

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(info.Name()) == ".json" {
			content, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", path, err)
				return err
			}

			stubs = append(stubs, json.RawMessage(content))
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking through directory: %v", err)
	}

	return stubs, nil
}
