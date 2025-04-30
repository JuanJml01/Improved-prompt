// Package main is the entry point for the tokinfo CLI application.
package main

import (
	"context" // Add context import
	"flag"
	"fmt"
	"log" // Using log for simple error reporting
	"os"  // Add os import

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
	flag.Parse()

	// --- Input Validation ---
	if *promptInput == "" {
		log.Fatal("Error: -p flag (prompt input) is required.") // Use log.Fatal for cleaner exit on error
	}

	fmt.Println("Starting Tokinfo: Prompt Enhancement Tool...") // Indicate start

	// --- Load Guidelines ---
	guidelines, err := config.LoadGuidelines("guidelines.json")
	if err != nil {
		log.Fatalf("Error loading guidelines: %v", err)
	}
	fmt.Println("Guidelines loaded.") // Progress message

	// --- Read User Prompt ---
	userPrompt, err := prompt.ReadInput(*promptInput)
	if err != nil {
		log.Fatalf("Error reading prompt input: %v", err)
	}
	fmt.Println("User prompt read successfully.") // Progress message

	// --- Initialize Gemini Client ---
	apiKey := os.Getenv("GEMINI_API_KEY") // Get API key from environment variable
	if apiKey == "" {
		log.Fatal("Error: GEMINI_API_KEY environment variable not set.")
	}

	// Create a context
	ctx := context.Background()

	// Initialize the Gemini client
	geminiClient, err := gemini.NewClient(ctx, apiKey)
	if err != nil {
		log.Fatalf("Error initializing Gemini client: %v", err)
	}
	defer geminiClient.Close()                // Ensure resources are released
	fmt.Println("Gemini client initialized.") // Progress message

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
	fmt.Println("Stage 1 analysis complete. Chosen technique:", analysisResult.ChosenTechniqueName)

	// --- User Interaction (Simulated/Placeholder) ---
	// For this structural setup, we assume no interactive step yet.
	// userAnswers := make(map[string]string) // Placeholder for answers if questions were asked
	// fmt.Println("Skipping user interaction step for now.")

	// --- Stage 2: Refinement ---
	chosenTechnique, found := config.GetTechniqueByName(guidelines.Techniques, analysisResult.ChosenTechniqueName)
	if !found {
		log.Fatalf("Error: Chosen technique '%s' not found in guidelines.", analysisResult.ChosenTechniqueName)
	}
	// For now, we use an empty map for user answers as the interaction step is skipped.
	userAnswers := make(map[string]string)
	enhancedPrompt, err := geminiClient.RefinePrompt(ctx, guidelines.Introduction, chosenTechnique.Complete, userPrompt, userAnswers)
	if err != nil {
		log.Fatalf("Error during Stage 2 Gemini call: %v", err)
	}
	fmt.Println("Stage 2 refinement complete.")

	// --- Output Enhanced Prompt ---
	err = prompt.HandleOutput(enhancedPrompt, *outputPath)
	if err != nil {
		log.Fatalf("Error handling output: %v", err)
	}

	// Final success message depends on output method
	if *outputPath != "" {
		fmt.Printf("Enhanced prompt successfully saved to %s\n", *outputPath)
	} else {
		fmt.Println("\n--- Enhanced Prompt ---")
		// The actual prompt content would be printed by HandleOutput
	}

}
