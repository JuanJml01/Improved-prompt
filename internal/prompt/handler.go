// Package prompt handles reading the user's prompt input (from string or file)
// and writing the final output (to stdout or file).
package prompt

import (
	"fmt"
	"os"
	"path/filepath" // Useful for checking extensions
	"strings"
)

// ReadInput determines if the input is a file path or a raw string,
// reads the content if it's a file, and returns the prompt string.
func ReadInput(inputPathOrString string, verbose bool) (string, error) {
	// Check for common text/markdown file extensions
	ext := strings.ToLower(filepath.Ext(inputPathOrString))
	isFilePath := ext == ".txt" || ext == ".md"

	// Alternative check: try to get file info, if it succeeds, it's likely a file.
	// _, err := os.Stat(inputPathOrString)
	// isFilePath := err == nil // If Stat succeeds without error, it's a file/dir

	if isFilePath {
		// Input is considered a file path
		if verbose {
			fmt.Printf("Reading prompt from file: %s\n", inputPathOrString)
		}
		content, err := os.ReadFile(inputPathOrString)
		if err != nil {
			return "", fmt.Errorf("failed to read prompt file '%s': %w", inputPathOrString, err)
		}
		return string(content), nil
	}

	// Input is considered a raw string
	if inputPathOrString == "" {
		return "", fmt.Errorf("prompt input cannot be empty")
	}
	return inputPathOrString, nil
}

// HandleOutput writes the provided content either to the specified outputPath file
// or to standard output if outputPath is empty.
func HandleOutput(content string, outputPath string, verbose bool) error {
	if outputPath == "" {
		// Write to standard output
		// The final prompt is always printed, so no verbose check here.
		_, err := fmt.Println(content) // fmt.Println writes to os.Stdout
		if err != nil {
			return fmt.Errorf("failed to write to standard output: %w", err)
		}
	} else {
		// Write to the specified file
		if verbose {
			fmt.Printf("Writing enhanced prompt to file: %s\n", outputPath)
		}
		// Ensure parent directories exist? (os.MkdirAll(filepath.Dir(outputPath), 0755)) - Consider adding
		err := os.WriteFile(outputPath, []byte(content), 0644) // Sensible default permissions
		if err != nil {
			return fmt.Errorf("failed to write output file '%s': %w", outputPath, err)
		}
	}
	return nil
}
