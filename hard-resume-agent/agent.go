package main

import (
	"context"
	"github.com/google/generative-ai-go/genai"
)

import _ "embed"

//go:embed data.txt
var kb string

type agent struct {
	client *genai.Client
}

func (a *agent) query(ctx context.Context, input string) (string, error) {
	model := a.client.GenerativeModel("gemini-pro")

	resp, err := model.GenerateContent(ctx, genai.Text(kb), genai.Text(input))
	if err != nil {
		return "", err
	}
	if len(resp.Candidates) == 0 {
		return "I am sorry, I did not understand", nil
	}
	p := resp.Candidates[0].Content.Parts[0]
	if txt, ok := p.(genai.Text); ok {
		return string(txt), nil
	}
	return "Returned another type", nil
}
