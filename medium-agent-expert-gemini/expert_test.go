package main

import (
	"context"
	"testing"
)

func TestEvaluate_LLMNotSet(t *testing.T) {
	ctx := context.Background()
	expert := &Expert{
		LLM:         nil,
		Prompt:      "What is your analysis?",
		Description: "An expert without an LLM.",
	}

	_, err := expert.Evaluate(ctx, "Test input")
	if err == nil || err.Error() != "LLM is not set" {
		t.Errorf("Expected error 'LLM is not set', got %v", err)
	}
}

func TestEvaluate_PromptNotSet(t *testing.T) {
	ctx := context.Background()
	expert := &Expert{
		LLM:         &MockLLM{},
		Prompt:      "",
		Description: "An expert without a prompt.",
	}

	_, err := expert.Evaluate(ctx, "Test input")
	if err == nil || err.Error() != "prompt is not set" {
		t.Errorf("Expected error 'prompt is not set', got %v", err)
	}
}

func TestEvaluate_EmptyResponse(t *testing.T) {
	ctx := context.Background()
	expert := &Expert{
		LLM:         &MockLLM{GenerateError: false, EmptyResponse: true},
		Prompt:      "What is your analysis?",
		Description: "An expert with an empty response.",
	}

	_, err := expert.Evaluate(ctx, "Test input")
	if err == nil || err.Error() != "empty response from LLM" {
		t.Errorf("Expected error 'empty response from LLM', got %v", err)
	}
}

func TestEvaluate_NilContext(t *testing.T) {
	ctx := context.TODO() // Using context.TODO() to simulate a nil context scenario
	expert := &Expert{
		LLM:         &MockLLM{},
		Prompt:      "What is your analysis?",
		Description: "An expert with a valid LLM and prompt.",
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Evaluate panicked with nil context: %v", r)
		}
	}()

	_, err := expert.Evaluate(ctx, "Test input")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestEvaluate_LLMError(t *testing.T) {
	ctx := context.Background()
	expert := &Expert{
		LLM:         &MockLLM{GenerateError: true, EmptyResponse: false},
		Prompt:      "What is your analysis?",
		Description: "An expert with an LLM that returns an error.",
	}

	_, err := expert.Evaluate(ctx, "Test input")
	if err == nil || err.Error() != "mock LLM error" {
		t.Errorf("Expected error 'mock LLM error', got %v", err)
	}
}
