# Tokinfo Go CLI Tool - Project Plan

This document outlines the planned structure and initial content for the `tokinfo` Go command-line tool.

## 1. Project Goal

To create the foundational Go project structure for `tokinfo`, a CLI tool designed to enhance user-provided prompts using the Gemini AI model based on prompt engineering guidelines stored in a JSON file. The focus is on the structure and commented placeholders, not the implementation logic.

## 2. Project Structure

The following directory structure will be used:

```
/home/juan/Documentos/go/Improved-prompt/
├── .gitignore         (exists)
├── LICENSE            (exists)
├── README.md          (exists)
├── main.go            # Entry point, CLI flags, workflow orchestration
├── guidelines.json    # Stores prompt engineering guidelines
└── internal/
    ├── config/
    │   └── config.go  # Logic for loading and parsing guidelines.json
    ├── gemini/
    │   └── client.go  # Logic for interacting with the Gemini API
    └── prompt/
        └── handler.go # Logic for reading prompt input & handling output
```

### Mermaid Diagram

```mermaid
graph TD
    A[/] --> B[main.go];
    A --> C[guidelines.json];
    A --> D[internal];
    D --> E[config];
    E --> F[config.go];
    D --> G[gemini];
    G --> H[client.go];
    D --> I[prompt];
    I --> J[handler.go];

    style A fill:#f9f,stroke:#333,stroke-width:2px
    style D fill:#ccf,stroke:#333,stroke-width:2px
    style E fill:#ccf,stroke:#333,stroke-width:2px
    style G fill:#ccf,stroke:#333,stroke-width:2px
    style I fill:#ccf,stroke:#333,stroke-width:2px
```

## 3. File Contents

### `guidelines.json`

```json
{
  "introduction": "Prompt engineering involves crafting effective inputs for AI models like Gemini to achieve desired outputs. This guide outlines key techniques.",
  "techniques": [
    {
      "name": "persona_pattern",
      "summarized": "Instruct the AI to adopt a specific persona (e.g., 'Act as a senior software engineer').",
      "complete": "Define a role for the AI. This helps set the context, tone, and expertise level for the response. Example: 'You are a helpful assistant specializing in explaining complex scientific concepts to a lay audience.'"
    },
    {
      "name": "few_shot_examples",
      "summarized": "Provide a few examples of the desired input/output format.",
      "complete": "Include 2-5 examples demonstrating the task you want the AI to perform. This is particularly effective for tasks involving specific formatting or style. Example: Input: 'Translate to French: Hello' Output: 'Bonjour'. Input: 'Translate to French: Goodbye' Output: 'Au revoir'."
    },
    {
      "name": "clear_instructions",
      "summarized": "State the task clearly and specifically.",
      "complete": "Avoid ambiguity. Clearly define the objective, the expected format of the output, and any constraints. Break down complex tasks into smaller, manageable steps if necessary."
    }
  ]
}
```

### `main.go`

```go
// Package main is the entry point for the tokinfo CLI application.
package main

import (
	"flag"
	"fmt"
	"log" // Using log for simple error reporting
	"os"

	// It's conventional to alias internal packages based on their directory name.
	// These imports will be uncommented as the packages are implemented.
	// config "tokinfo/internal/config"
	// gemini "tokinfo/internal/gemini"
	// prompt "tokinfo/internal/prompt"
)

func main() {
	// Define command-line flags for user input and output options.
	promptInput := flag.String("p", "", "Prompt string or path to prompt file (.txt, .md) (required)")
	outputPath := flag.String("g", "", "Optional path to save the generated prompt")
	flag.Parse()

	// --- Input Validation ---
	if *promptInput == "" {
		log.Fatal("Error: -p flag (prompt input) is required.") // Use log.Fatal for cleaner exit on error
	}

	fmt.Println("Starting Tokinfo: Prompt Enhancement Tool...") // Indicate start

	// --- Load Guidelines ---
	// guidelines, err := config.LoadGuidelines("guidelines.json")
	// if err != nil {
	// 	log.Fatalf("Error loading guidelines: %v", err)
	// }
	// fmt.Println("Guidelines loaded.") // Progress message

	// --- Read User Prompt ---
	// userPrompt, err := prompt.ReadInput(*promptInput)
	// if err != nil {
	// 	log.Fatalf("Error reading prompt input: %v", err)
	// }
	// fmt.Println("User prompt read successfully.") // Progress message

	// --- Initialize Gemini Client ---
	// apiKey := os.Getenv("GEMINI_API_KEY") // Example: Get API key from environment variable
	// if apiKey == "" {
	// 	log.Fatal("Error: GEMINI_API_KEY environment variable not set.")
	// }
	// geminiClient, err := gemini.NewClient(apiKey)
	// if err != nil {
	// 	log.Fatalf("Error initializing Gemini client: %v", err)
	// }
	// defer geminiClient.Close() // Ensure resources are released
	// fmt.Println("Gemini client initialized.") // Progress message

	// --- Stage 1: Analysis & Clarification ---
	// analysisResult, err := geminiClient.AnalyzePrompt(guidelines.Introduction, guidelines.Techniques, userPrompt)
	// if err != nil {
	// 	log.Fatalf("Error during Stage 1 Gemini call: %v", err)
	// }
	// fmt.Println("Stage 1 analysis complete. Chosen technique:", analysisResult.ChosenTechniqueName)

	// --- User Interaction (Simulated/Placeholder) ---
	// For this structural setup, we assume no interactive step yet.
	// userAnswers := make(map[string]string) // Placeholder for answers if questions were asked
	// fmt.Println("Skipping user interaction step for now.")

	// --- Stage 2: Refinement ---
	// chosenTechnique, found := config.GetTechniqueByName(guidelines.Techniques, analysisResult.ChosenTechniqueName)
	// if !found {
	//     log.Fatalf("Error: Chosen technique '%s' not found in guidelines.", analysisResult.ChosenTechniqueName)
	// }
	// enhancedPrompt, err := geminiClient.RefinePrompt(guidelines.Introduction, chosenTechnique.Complete, userPrompt, userAnswers)
	// if err != nil {
	// 	log.Fatalf("Error during Stage 2 Gemini call: %v", err)
	// }
	// fmt.Println("Stage 2 refinement complete.")

	// --- Output Enhanced Prompt ---
	// err = prompt.HandleOutput(enhancedPrompt, *outputPath)
	// if err != nil {
	// 	log.Fatalf("Error handling output: %v", err)
	// }

	// Final success message depends on output method
	// if *outputPath != "" {
	// 	fmt.Printf("Enhanced prompt successfully saved to %s\n", *outputPath)
	// } else {
	// 	fmt.Println("\n--- Enhanced Prompt ---")
	// 	// The actual prompt content would be printed by HandleOutput
	// }

	// --- Temporary Placeholder Output ---
	fmt.Println("\nPlaceholder: Workflow completed (actual logic pending).")
	fmt.Printf("Input Prompt Flag (-p): %s\n", *promptInput)
	fmt.Printf("Output Path Flag (-g): %s\n", *outputPath)
	// End Temporary Placeholder Output ---
}
```

### `internal/config/config.go`

```go
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
```

### `internal/gemini/client.go`

```go
// Package gemini provides the client for interacting with the Google Gemini API.
package gemini

import (
	"context" // Gemini client likely requires context
	"fmt"     // For error formatting

	// Official Gemini Go client package import path needs to be added here.
	// Example: "google.golang.org/api/option"
	// Example: "github.com/google/generative-ai-go/genai"
)

// Client wraps the official Gemini client and provides specific methods for tokinfo.
type Client struct {
	// internalClient *genai.Client // Example: Store the official client instance
	apiKey string // Store API key if needed for methods
}

// AnalysisResult holds the structured data returned from the Stage 1 analysis call.
type AnalysisResult struct {
	ChosenTechniqueName string
	ClarifyingQuestions []string // Questions generated by Gemini for the user
	// Add any other relevant fields from the analysis, e.g., Gemini's reasoning.
}

// NewClient initializes and returns a new Gemini client wrapper.
// It requires the API key for authentication.
func NewClient(apiKey string) (*Client, error) {
	// Implementation details:
	// 1. Use the official genai package to create a new client instance.
	//    (e.g., genai.NewClient(context.Background(), option.WithAPIKey(apiKey)))
	// 2. Handle potential initialization errors.
	// 3. Return a new instance of our wrapper Client struct.

	if apiKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}
	// Placeholder implementation
	// officialClient, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	// if err != nil {
	//     return nil, fmt.Errorf("failed to create genai client: %w", err)
	// }

	return &Client{apiKey: apiKey /*, internalClient: officialClient*/}, nil // Placeholder return
}

// Close releases any resources held by the client.
func (c *Client) Close() error {
	// Implementation details:
	// 1. Call the Close method on the underlying official client if it exists.
	//    (e.g., return c.internalClient.Close())
	// 2. Handle potential errors during closing.
	fmt.Println("Gemini client resources released (placeholder).") // Placeholder action
	return nil
}

// AnalyzePrompt performs the Stage 1 interaction with the Gemini API.
// It sends the context and user prompt, requesting analysis and clarifying questions.
func (c *Client) AnalyzePrompt(ctx context.Context, intro string, summarizedTechniques string, userPrompt string) (*AnalysisResult, error) {
	// Implementation details:
	// 1. Construct the combined prompt for the Gemini API based on inputs.
	// 2. Use the c.internalClient to send the request (e.g., GenerateContent).
	// 3. Parse the response to extract the chosen technique name and questions.
	// 4. Handle API errors and response parsing errors.

	// Placeholder implementation
	fmt.Printf("Simulating Stage 1 API call for prompt: %s...\n", userPrompt)
	// Simulate finding a technique and asking a question
	if len(userPrompt) > 10 { // Arbitrary condition for simulation
		return &AnalysisResult{
			ChosenTechniqueName: "clear_instructions", // Example
			ClarifyingQuestions: []string{"What is the desired output format? (e.g., JSON, bullet points)"},
		}, nil
	}
	return &AnalysisResult{ChosenTechniqueName: "persona_pattern", ClarifyingQuestions: []string{}}, nil // Placeholder return
}

// RefinePrompt performs the Stage 2 interaction with the Gemini API.
// It sends the context, chosen technique details, original prompt, and any user answers
// to generate the final enhanced prompt.
func (c *Client) RefinePrompt(ctx context.Context, intro string, completeTechniqueDesc string, userPrompt string, answers map[string]string) (string, error) {
	// Implementation details:
	// 1. Construct the combined prompt for the Gemini API, incorporating all inputs.
	// 2. Use the c.internalClient to send the request.
	// 3. Parse the response to extract the final enhanced prompt string.
	// 4. Handle API errors and response parsing errors.

	// Placeholder implementation
	fmt.Printf("Simulating Stage 2 API call to refine prompt: %s...\n", userPrompt)
	refined := fmt.Sprintf("Enhanced version of '%s' using technique description: '%s'. User answers: %v", userPrompt, completeTechniqueDesc, answers)
	return refined, nil // Placeholder return
}

```

### `internal/prompt/handler.go`

```go
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
func ReadInput(inputPathOrString string) (string, error) {
	// Check for common text/markdown file extensions
	ext := strings.ToLower(filepath.Ext(inputPathOrString))
	isFilePath := ext == ".txt" || ext == ".md"

	// Alternative check: try to get file info, if it succeeds, it's likely a file.
	// _, err := os.Stat(inputPathOrString)
	// isFilePath := err == nil // If Stat succeeds without error, it's a file/dir

	if isFilePath {
		// Input is considered a file path
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
func HandleOutput(content string, outputPath string) error {
	if outputPath == "" {
		// Write to standard output
		_, err := fmt.Println(content) // fmt.Println writes to os.Stdout
		if err != nil {
			return fmt.Errorf("failed to write to standard output: %w", err)
		}
	} else {
		// Write to the specified file
		// Ensure parent directories exist? (os.MkdirAll(filepath.Dir(outputPath), 0755)) - Consider adding
		err := os.WriteFile(outputPath, []byte(content), 0644) // Sensible default permissions
		if err != nil {
			return fmt.Errorf("failed to write output file '%s': %w", outputPath, err)
		}
	}
	return nil
}
```

## 4. Next Steps

This plan can be used as a reference for a developer (or an AI in Code mode) to implement the `tokinfo` tool. The next logical step would be to switch to a mode capable of writing `.go` and `.json` files (like Code mode) and use this plan to create the actual project files.