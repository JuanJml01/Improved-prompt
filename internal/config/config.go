// Package config handles loading and accessing the prompt engineering guidelines.
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Technique defines the structure for a single prompt engineering technique.
type Technique struct {
	Name       string `json:"name"`
	Summarized string `json:"summarized"`
	Complete   string `json:"complete"`
}

// Guidelines defines the overall structure of the guidelines JSON file.
type Guidelines struct {
	Introduction string      `json:"introduction"`
	Techniques   []Technique `json:"techniques"`
}

// LoadGuidelines reads the specified JSON file and parses it into a Guidelines struct.
// It returns the populated struct or an error if reading/parsing fails.
func LoadGuidelines(filePath string) (*Guidelines, error) {
	// Implementation details:
	// 1. Use os.ReadFile to read the file content.
	// 2. Use json.Unmarshal to parse the content into the Guidelines struct.
	// 3. Return the struct and nil error, or nil and the encountered error.

	// Placeholder implementation (returns error until implemented)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read guidelines file '%s': %w", filePath, err)
	}

	var guidelines Guidelines
	err = json.Unmarshal(data, &guidelines)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal guidelines JSON from '%s': %w", filePath, err)
	}

	// Basic validation example (can be expanded)
	if guidelines.Introduction == "" || len(guidelines.Techniques) == 0 {
		return nil, fmt.Errorf("guidelines file '%s' is missing introduction or techniques", filePath)
	}

	return &guidelines, nil // Placeholder return
}

// GetTechniqueByName searches the list of techniques for one matching the given name.
// It returns the technique and true if found, otherwise nil and false.
func GetTechniqueByName(techniques []Technique, name string) (*Technique, bool) {
	for i := range techniques {
		if techniques[i].Name == name {
			return &techniques[i], true // Return pointer to the technique in the slice
		}
	}
	return nil, false
}
