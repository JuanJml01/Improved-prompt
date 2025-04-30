package gemini

// ResponseSchema defines the structure for the Gemini API response schema.
// It corresponds to the JSON schema provided for prompt enhancement.
type ResponseSchema struct {
	ChoseTechnique      string   `json:"ChoseTechnique"`
	ClarifyingQuestions []string `json:"ClarifyingQuestions"`
}
