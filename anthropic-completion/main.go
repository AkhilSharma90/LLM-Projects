package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

func main() {
	// Check if API key is set
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		log.Fatal("ANTHROPIC_API_KEY environment variable not set")
	}

	// Initialize the Anthropic client (API key will be read from env var)
	llm, err := anthropic.New(
		anthropic.WithModel("claude-3-sonnet-20240229"), // Or another model
	)
	if err != nil {
		log.Fatalf("Error creating Anthropic client: %v", err)
	}

	ctx := context.Background()

	// Generate completion with streaming
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, "Hi Claude, write a para about Golang powered AI systems",
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		log.Fatalf("Error generating completion: %v", err)
	}

	fmt.Println("\nCompletion:", completion)
}
