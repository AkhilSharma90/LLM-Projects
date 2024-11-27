package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
)

type ExpertInterface interface {
	Evaluate(ctx context.Context, input string) (*ExpertResult, error)
}

type Expert struct {
	ExpertInterface

	LLM         llms.Model
	CallOptions *llms.CallOptions `json:"options"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Prompt      string            `json:"prompt"`
}

type ExpertResult struct {
	Response *llms.ContentChoice `json:"response"`
	Expert   *Expert             `json:"expert"`
	Text     string              `json:"text"`
	Err      error               `json:"err"`
}

// Evaluate processes the input string using the configured language model (LLM)
// and returns an ExpertResult containing the response from the LLM.
//
// Parameters:
//
//	ctx - The context for controlling cancellation and timeouts.
//	input - The input string to be evaluated by the LLM.
//
// Returns:
//
//	*ExpertResult - A struct containing the response from the LLM, the expert
//	                instance, the text of the response, and any error encountered.
//	error - An error if the LLM is not set, the prompt is not set, or if there
//	        is an issue generating content from the LLM.
func (e *Expert) Evaluate(ctx context.Context, input string) (*ExpertResult, error) {

	if e.LLM == nil {
		return nil, fmt.Errorf("LLM is not set")
	}

	if e.Prompt == "" {
		return nil, fmt.Errorf("prompt is not set")
	}

	// Imbue the personality
	systemMessage := llms.MessageContent{
		Role:  llms.ChatMessageTypeSystem,
		Parts: []llms.ContentPart{llms.TextPart(e.Description)},
	}

	// Rephrase the question
	newPrompt := fmt.Sprintf("%s\n%s", e.Prompt, input)

	// Create a user message
	userMessage := llms.MessageContent{
		Role:  llms.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{llms.TextPart(newPrompt)},
	}

	// And ask the AI!
	content, err := e.LLM.GenerateContent(ctx, []llms.MessageContent{systemMessage, userMessage})

	if err != nil {
		return nil, err
	}

	if content == nil || len(content.Choices) == 0 {
		return nil, fmt.Errorf("empty response from LLM")
	}

	return &ExpertResult{
		Expert:   e,
		Text:     content.Choices[0].Content,
		Response: content.Choices[0],
		Err:      err,
	}, nil
}
