// Set the GOOGLE_API_KEY env var to your API key taken from https://aistudio.google.com/apikey
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GOOGLE_API_KEY")
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	prompt := "Who was the second person to walk on the moon?"
	answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}
