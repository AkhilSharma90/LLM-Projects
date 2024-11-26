# Google AI LLM Generation Example with LangChainGo

## Overview

A simple Go application demonstrating how to use the Google AI Language Model via LangChainGo for text generation. This minimal example shows basic prompt-based text generation using Google's AI platform.

## Prerequisites

- Go (version 1.20+)
- Google AI Studio API Key

## Getting API Key

1. Visit [Google AI Studio](https://aistudio.google.com/apikey)
2. Create and copy your API key
3. Set as an environment variable

## Dependencies

```bash
go get github.com/tmc/langchaingo/llms
go get github.com/tmc/langchaingo/llms/googleai
```

## Environment Setup

Set your Google AI API key:

```bash
# macOS/Linux
export GOOGLE_API_KEY='your_google_ai_api_key_here'

# Windows PowerShell
$env:GOOGLE_API_KEY='your_google_ai_api_key_here'
```

## Project Structure

- `main.go`: Single-file application
  - Demonstrates simple text generation
  - Uses Google AI through LangChainGo

## Key Features

- Simple prompt-based text generation
- Uses Google AI language model
- Minimal, straightforward implementation
- Environment-based API key configuration

## How to Run

```bash
go mod tidy
go run main.go
```

## Customization Options

- Modify input prompt
- Experiment with different queries
- Adjust generation parameters (if needed)

## Example Use Cases

- Answering historical questions
- Generating creative text
- Exploring AI-generated content
- Quick information retrieval

## Potential Improvements

- Add error handling
- Implement more complex prompt strategies
- Support multiple generation configurations
- Add logging and metrics

## Technologies

- **Language**: Go
- **AI**: Google AI
- **Library**: LangChainGo

## Limitations

- Requires active Google AI API key
- Dependent on Google AI model capabilities

## Error Handling

- Checks API key configuration
- Handles generation errors
- Provides clear error messages

## Typical Output Example

```
Buzz Aldrin
```

## Troubleshooting

- Verify API key is correctly set
- Check internet connectivity
- Ensure latest dependencies are installed