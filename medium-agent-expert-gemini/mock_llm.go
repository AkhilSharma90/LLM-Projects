package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
)

type MockLLM struct {
	EmptyResponse bool
	GenerateError bool
}

func (m *MockLLM) GenerateContent(_ context.Context, _ []llms.MessageContent, _ ...llms.CallOption) (*llms.ContentResponse, error) {

	if m.GenerateError {
		return nil, fmt.Errorf("mock LLM error")
	}

	if m.EmptyResponse {
		return nil, fmt.Errorf("empty response from LLM")
	}

	choice := &llms.ContentChoice{Content: "Mock response"}
	return &llms.ContentResponse{Choices: []*llms.ContentChoice{choice}}, nil
}

func (m *MockLLM) Call(_ context.Context, _ string, _ ...llms.CallOption) (string, error) {
	return "Mock response", nil
}
