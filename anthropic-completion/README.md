# Golang Anthropic LLM Streaming Example

## Overview

This is a simple Go project demonstrating how to use the Anthropic Claude AI model with LangChainGo for generating text completions with streaming output. The example shows how to:

- Initialize an Anthropic LLM client
- Generate text completions
- Stream the AI's response in real-time

## Prerequisites

- Go (version 1.20 or higher)
- Anthropic API Key

## Setup

1. **Install Dependencies**:

   ```bash
   go mod init anthropic-llm-example
   go get github.com/tmc/langchaingo/llms
   go get github.com/tmc/langchaingo/llms/anthropic
   ```

2. **API Key Configuration**:

   You must set your Anthropic API key as an environment variable:

   ```bash
   export ANTHROPIC_API_KEY='your_api_key_here'
   ```

## Running the Project

```bash
go run main.go
```

## How It Works

The script does the following:

1. Checks if the Anthropic API key is set
2. Creates an Anthropic LLM client with Claude-3-Sonnet model
3. Generates a text completion with streaming output
4. Prints the streamed response in real-time

## Customization

- Change the model by modifying `anthropic.WithModel()`.
- Adjust the temperature or prompt as needed.
- Add more complex prompt engineering.

## Error Handling

The script includes basic error handling:

- Checks for missing API key
- Handles client initialization errors
- Manages completion generation errors

## Dependencies

- `github.com/tmc/langchaingo/llms`
- `github.com/tmc/langchaingo/llms/anthropic`