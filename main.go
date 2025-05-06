// Package main is the entry point for the tokinfo CLI application.
package main

import (
	"bufio"
	"context" // Add context import
	"flag"
	"fmt"
	"log" // Using log for simple error reporting
	"os"  // Add os import
	"strings"

	// It's conventional to alias internal packages based on their directory name.
	// These imports will be uncommented as the packages are implemented.
	config "tokinfo/internal/config"
	gemini "tokinfo/internal/gemini"
	prompt "tokinfo/internal/prompt"
)

func main() {
	// Define command-line flags for user input and output options.
	promptInput := flag.String("p", "", "Prompt string or path to prompt file (.txt, .md) (required)")
	// outputPath is currently unused in the Analysis & Clarification phase,
	// but is kept for future phases.
	outputPath := flag.String("g", "", "Optional path to save the generated prompt")
	verbose := flag.Bool("verbose", false, "Enable verbose output") // Add verbose flag
	flag.Parse()

	// --- Input Validation ---
	if *promptInput == "" {
		log.Fatal("Error: -p flag (prompt input) is required.") // Use log.Fatal for cleaner exit on error
	}

	if *verbose {
		fmt.Println("Starting Tokinfo: Prompt Enhancement Tool...") // Indicate start
	}

	// --- Load Guidelines ---
	guidelines, err := config.LoadGuidelines("guidelines.json", *verbose) // Pass verbose flag
	if err != nil {
		log.Fatalf("Error loading guidelines: %v", err)
	}
	if *verbose {
		fmt.Println("Guidelines loaded.") // Progress message
	}

	// --- Read User Prompt ---
	userPrompt, err := prompt.ReadInput(*promptInput, *verbose) // Pass verbose flag
	if err != nil {
		log.Fatalf("Error reading prompt input: %v", err)
	}
	if *verbose {
		fmt.Println("User prompt read successfully.") // Progress message
	}

	// --- Initialize Gemini Client ---
	apiKey := os.Getenv("GEMINI_API_KEY") // Get API key from environment variable
	if apiKey == "" {
		log.Fatal("Error: GEMINI_API_KEY environment variable not set.")
	}

	// Create a context
	ctx := context.Background()

	// Initialize the Gemini client
	geminiClient, err := gemini.NewClient(ctx, apiKey, *verbose) // Pass verbose flag
	if err != nil {
		log.Fatalf("Error initializing Gemini client: %v", err)
	}
	defer geminiClient.Close() // Ensure resources are released
	if *verbose {
		fmt.Println("Gemini client initialized.") // Progress message
	}

	// --- Stage 1: Analysis & Clarification ---
	// Summarize techniques for the Gemini API call
	var summarizedTechniques string
	for _, tech := range guidelines.Techniques {
		summarizedTechniques += fmt.Sprintf("- %s: %s\n", tech.Name, tech.Summarized)
	}

	// --- Stage 1: Analysis & Clarification ---
	// Call the AnalyzePrompt method on the Gemini client.
	// This sends the introduction, summarized techniques, and user prompt to the Gemini model
	// for analysis and to get clarifying questions.
	analysisResult, err := geminiClient.AnalyzePrompt(ctx, guidelines.Introduction, summarizedTechniques, userPrompt)
	if err != nil {
		log.Fatalf("Error during Stage 1 Gemini call: %v", err)
	}
	if *verbose {
		fmt.Println("Stage 1 analysis complete. Chosen technique:", analysisResult.ChosenTechniqueName)
	}

	// --- User Interaction ---
	userAnswers := make(map[string]string) // Initialize map for answers
	if len(analysisResult.ClarifyingQuestions) > 0 {
		if *verbose {
			fmt.Println("\nPlease answer the following questions to help refine the prompt:")
		}
		reader := bufio.NewReader(os.Stdin) // Create a reader for input

		// Iterate over the slice of question strings
		for i, questionText := range analysisResult.ClarifyingQuestions {
			fmt.Printf("- %s: ", questionText) // Print the question text
			answer, err := reader.ReadString('\n')
			if err != nil {
				// Basic error handling for reading input
				log.Printf("Warning: Could not read answer for question %d ('%s'): %v. Skipping.", i+1, questionText, err)
				continue // Skip this question if reading fails
			}
			// Trim newline characters (\r\n on Windows, \n on Unix)
			// Use the question text as the key for the answer map
			userAnswers[questionText] = strings.TrimSpace(answer)
		}
		if *verbose {
			fmt.Println("Thank you for your answers.")
		}
	} else {
		if *verbose {
			fmt.Println("No clarifying questions needed based on the analysis.")
		}
	}

	// --- Stage 2: Refinement ---
	chosenTechnique, found := config.GetTechniqueByName(guidelines.Techniques, analysisResult.ChosenTechniqueName)
	if !found {
		log.Fatalf("Error: Chosen technique '%s' not found in guidelines.", analysisResult.ChosenTechniqueName)
	}
	// userAnswers map is now populated from the interaction step above (if any questions were asked).
	enhancedPrompt, err := geminiClient.RefinePrompt(ctx, guidelines.Introduction, chosenTechnique.Complete, userPrompt, userAnswers)
	if err != nil {
		log.Fatalf("Error during Stage 2 Gemini call: %v", err)
	}
	if *verbose {
		fmt.Println("Stage 2 refinement complete.")
	}

	// --- Stage 3: Execute Enhanced Prompt ---
	if *verbose {
		fmt.Println("\nExecuting enhanced prompt with Gemini...")
	}

	// The final result should ALWAYS be printed, regardless of verbose flag.
	fmt.Println(enhancedPrompt)

	// Final success message depends on output method
	if *verbose && *outputPath != "" {
		fmt.Printf("Enhanced prompt successfully saved to %s\n", *outputPath)
	}
	// No need to print the enhanced prompt again here if HandleOutput already did
	// The final Gemini response is printed above, outside the verbose check.

}

// Helper function to access refineConfig (needs to be added to client.go or accessed differently)
// For now, let's assume we need to add a getter in client.go
// Alternatively, pass nil if GenerateResponse handles it:
// finalResult, err := geminiClient.GenerateResponse(ctx, "gemini-2.0-flash", enhancedPrompt, nil)
