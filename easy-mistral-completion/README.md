# LangChainGo Mistral AI Text Generation Example

## Overview

This Go application demonstrates text generation using the LangChainGo library with Mistral AI's language models. The example showcases two key approaches to text generation:

1. Streaming text generation (real-time output)
2. Standard text generation with model-specific parameters

## Features

- Integration with Mistral AI language models
- Streaming text generation
- Configurable generation parameters
- Multiple model support
- Flexible prompt handling

## Prerequisites

- Go (version 1.16 or later)
- Mistral AI API Key

## Installation

1. Clone the repository:

   ```bash
   git clone <your-repository-url>
   cd <project-directory>
   ```

2. Set up your Mistral AI API key as an environment variable:

   ```bash
   export MISTRAL_API_KEY='your-api-key-here'
   ```

## Dependencies

This project uses the following key dependencies:
- `github.com/tmc/langchaingo/llms`
- `github.com/tmc/langchaingo/llms/mistral`

## Dependency Installation

Install dependencies using:

```bash
go mod init your-project-name
go get github.com/tmc/langchaingo/llms
go get github.com/tmc/langchaingo/llms/mistral
```

## Key Concepts Demonstrated

### 1. Streaming Text Generation

- Uses `llms.WithStreamingFunc()` to provide real-time output
- Allows watching text generation as it happens
- Useful for interactive applications or progress tracking

### 2. Standard Text Generation

- Uses `llms.GenerateFromSinglePrompt()` with fixed parameters
- Demonstrates model selection
- Shows how to control generation creativity via temperature

## Usage Examples

### Streaming Generation

The first generation demonstrates streaming:
- Prompt: "Who invented telescope?"
- Model: "open-mistral-7b"
- Temperature: 0.8 (more creative responses)

### Standard Generation

The second generation shows standard generation:
- Prompt: "Who was the first man to go to space?"
- Model: "mistral-small-latest"
- Temperature: 0.2 (more focused responses)

## Running the Project

```bash
go run main.go
```

## Configuration Options

You can customize text generation by modifying:
- Prompt content
- Model selection
- Temperature (creativity level)
- Streaming vs. non-streaming generation

## Error Handling

The application includes basic error handling:
- Checks for errors when initializing the Mistral AI client
- Handles potential errors during content generation