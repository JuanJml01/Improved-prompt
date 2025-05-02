// Package gemini provides the client for interacting with the Google Gemini API.
package gemini

import (
	"context"       // Gemini client likely requires context
	"encoding/json" // For JSON parsing
	"fmt"           // For error formatting

	// Official Gemini Go client package import path needs to be added here.
	// Example: "google.golang.org/api/option"
	// Example: "github.com/google/generative-ai-go/genai"
	"google.golang.org/genai"
)

// Client wraps the official Gemini client and provides specific methods for tokinfo.
type Client struct {
	*genai.Client // Embed the official client
	analyzeConfig *genai.GenerateContentConfig
	refineConfig  *genai.GenerateContentConfig
}

// AnalysisResult holds the structured data returned from the Stage 1 analysis call.
type AnalysisResult struct {
	ChosenTechniqueName string   `json:"ChoseTechnique"`      // Match JSON key "ChoseTechnique"
	ClarifyingQuestions []string `json:"ClarifyingQuestions"` // Match JSON key "ClarifyingQuestions"
	// Add any other relevant fields from the analysis, e.g., Gemini's reasoning.
}

// NewClient initializes and returns a new Gemini client wrapper.
// It requires the API key for authentication.
func NewClient(ctx context.Context, apiKey string) (*Client, error) {
	// Use the official genai package to create a new client instance.
	// Handle potential initialization errors.
	// Return a new instance of our wrapper Client struct.

	if apiKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}

	// Create ClientConfig
	cfg := &genai.ClientConfig{
		APIKey: apiKey,
		// Add other config options if needed, e.g., Backend, Project, Location
	}

	// Create the official client
	officialClient, err := genai.NewClient(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %w", err)
	}

	// Define the GenerateContentConfig for the AnalyzePrompt function, using the schema.
	analyzeConfig := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"ChoseTechnique": {Type: genai.TypeString},
				"ClarifyingQuestions": {
					Type:  genai.TypeArray,
					Items: &genai.Schema{Type: genai.TypeString},
				},
			},
		},
	}

	// Define a simple GenerateContentConfig for the RefinePrompt function.
	// This config does not require a specific schema.
	refineConfig := &genai.GenerateContentConfig{
		// Add any simple configurations needed for refinement, or leave empty.
		// For now, we'll keep it minimal as per the request for a "simple" config.
	}

	// Return our wrapper client embedding the official client and the configs
	return &Client{
		Client:        officialClient,
		analyzeConfig: analyzeConfig,
		refineConfig:  refineConfig,
	}, nil
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
// It uses the analyzeConfig with the defined schema for structured output.
func (c *Client) AnalyzePrompt(ctx context.Context, intro string, summarizedTechniques string, userPrompt string) (*AnalysisResult, error) {
	// Construct the combined prompt for the Gemini API based on inputs.
	prompt := fmt.Sprintf(`Prompt Engineering Guide:
%s

User’s Raw Prompt:
%s

Task:
Using only the techniques described in the Prompt Engineering Guide, analyze the User’s Raw Prompt and decide:

1. Which single prompt-engineering technique you will apply.
2. What clarifying questions (if any) you need to ask before rewriting it — and for each question, provide an example of an appropriate answer.

Output:
Respond with exactly this JSON schema—no extra keys or prose:

{
	 "type": "object",
	 "properties": {
	   "ChoseTechnique": {
	     "type": "string",
	     "description": "The name of the chosen technique from the Guide."
	   },
	   "ClarifyingQuestions": {
	     "type": "array",
	     "items": {
	       "type": "object",
	       "properties": {
	         "question": {
	           "type": "string",
	           "description": "The clarifying question to ask the user."
	         },
	         "exampleAnswer": {
	           "type": "string",
	           "description": "A sample answer that the user might give."
	         }
	       },
	       "required": ["question", "exampleAnswer"]
	     }
	   }
	 },
	 "required": ["ChoseTechnique", "ClarifyingQuestions"]
}`, intro+"\n\n"+summarizedTechniques, userPrompt)

	// Use the GenerateResponse helper function with the analyzeConfig.
	generatedText, err := c.GenerateResponse(ctx, "gemini-2.0-flash", prompt, c.analyzeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content for analysis: %w", err)
	}

	// Unmarshal the JSON response text into the AnalysisResult struct.
	var result AnalysisResult
	err = json.Unmarshal([]byte(generatedText), &result)
	if err != nil {
		// Log the problematic JSON for debugging if needed
		// fmt.Printf("Failed to unmarshal JSON: %s\n", generatedText)
		return nil, fmt.Errorf("failed to unmarshal analysis response JSON: %w", err)
	}

	// Return the parsed result.
	return &result, nil
}

// RefinePrompt performs the Stage 2 interaction with the Gemini API.
// It sends the context, chosen technique details, original prompt, and any user answers
// to generate the final enhanced prompt.
// It uses the simple refineConfig.
func (c *Client) RefinePrompt(ctx context.Context, intro string, completeTechniqueDesc string, userPrompt string, answers map[string]string) (string, error) {
	// Construct the combined prompt for the Gemini API, incorporating all inputs.
	prompt := fmt.Sprintf(`%s  (intro)
%s (completeTechniqueDesc)
%s  (userprompt)
%v (answers)

You are a prompt enhancement tool that rigorously applies the provided engineering guidelines. Refine the user's original "{prompt}" by:
1. **Integrating** the context from:
	  - {intro} (core principles)
	  - {technique description} (methodology)
	  - {extra information} (additional constraints/requirements)
2. **Enhancing** specificity, structure, and clarity while **preserving every element** of the original prompt.
3. **Formatting** the output as a standalone, optimized prompt in English with no explanations, headers, or markdown.

**Constraints:**
- Do **not** add, remove, or reinterpret concepts from "{prompt}".
- Use **only** the context from {intro}, {technique description}, and {extra information}.
- Output **exclusively** the final enhanced prompt.

**Example Transformation:**
Original: "Explain blockchain"
Enhanced: "Describe blockchain technology in 3 steps using a baking analogy for non-technical audiences. Highlight decentralization and security. Avoid cryptocurrency mentions."`,
		intro, completeTechniqueDesc, userPrompt, answers,
	)

	// Use the GenerateResponse helper function with the refineConfig.
	refinedPrompt, err := c.GenerateResponse(ctx, "gemini-2.0-flash", prompt, c.refineConfig)
	if err != nil {
		return "", fmt.Errorf("failed to generate content for refinement: %w", err)
	}

	// Assuming the response is plain text, return the generated text.
	return refinedPrompt, nil
}

// GenerateResponse calls the Gemini API's GenerateContent method to get a response.
// It takes the context, model name, prompt, and configuration, and returns the generated text or an error.
func (c *Client) GenerateResponse(ctx context.Context, modelName string, prompt string, config *genai.GenerateContentConfig) (string, error) {
	// Call the embedded genai.Client's GenerateContent method
	result, err := c.Client.Models.GenerateContent(ctx, modelName, genai.Text(prompt), config)
	if err != nil {
		// Handle the error from the API call
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	// Extract the text from the result
	generatedText := result.Text()

	// Return the extracted text and nil error
	return generatedText, nil
}
