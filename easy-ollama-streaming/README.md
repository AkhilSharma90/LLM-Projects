# LangChainGo Ollama Local LLM Streaming Example

## Overview

This Go application demonstrates using LangChainGo with Ollama to generate a creative company name for Go-backed LLM tools. It showcases local AI model interaction, streaming generation, and context-based content generation.

## Features

- Local AI model generation with Ollama
- Streaming text output
- Context-aware content generation
- Flexible system and user message configuration

## Prerequisites

- Go (version 1.16 or later)
- [Ollama](https://ollama.com/) installed and running
- Ollama model downloaded (e.g., Gemma 2B)

## Installation

1. Install Ollama

   ```bash
   # Follow instructions at https://ollama.com/
   curl https://ollama.ai/install.sh | sh
   ```

2. Pull the Gemma model

   ```bash
   ollama pull gemma:2b
   ```

3. Clone the repository:

   ```bash
   git clone <your-repository-url>
   cd <project-directory>
   ```

## Dependencies

This project uses the following key dependencies:
- `github.com/tmc/langchaingo/llms`
- `github.com/tmc/langchaingo/llms/ollama`

## Dependency Installation

Install dependencies using:

```bash
go mod init your-project-name
go get github.com/tmc/langchaingo/llms
go get github.com/tmc/langchaingo/llms/ollama
```

## Key Concepts Demonstrated

### Local AI Model Generation

- Uses Ollama for local, private AI model inference
- Supports various lightweight models like Gemma 2B
- Enables AI-powered tasks without external API calls

### Streaming Text Generation

- Real-time output of generated content
- Demonstrates interactive AI response generation
- Useful for creating responsive AI interfaces

### Context-Based Generation

- Uses system and user messages to guide AI response
- Provides context and role-setting for more targeted outputs

## Usage

### Running the Project

Ensure Ollama is running, then execute:

```bash
go run main.go
```

### Customization Options

You can modify:
- Ollama model (currently set to `gemma:2b`)
- System message context
- Prompt for name generation
- Streaming output handling

## Model Selection

The example uses `gemma:2b`, but you can experiment with other Ollama models:
- `llama2:7b`
- `mistral:7b`
- `phi:2.7b`

To change the model, modify the `ollama.WithModel()` parameter.

## Error Handling

The application includes basic error handling:
- Checks for Ollama model initialization errors
- Manages content generation errors

## Performance Considerations

- Smaller models (2B-7B) are faster but may provide less nuanced responses
- Larger models offer more sophisticated outputs at the cost of slower generation