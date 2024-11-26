# Golang Example using Groq and Llama3

## Overview

A Go application demonstrating AI-powered poetry generation using:
- Groq Cloud Platform
- Llama3-8b Language Model
- LangChainGo for AI interaction
- Streaming text generation

## Prerequisites

- Go (version 1.20+)
- Groq Cloud Account

## Getting Started

### 1. Groq API Key

1. Visit [Groq Console](https://console.groq.com/keys)
2. Create and copy your API key

## Dependencies

```bash
go get github.com/tmc/langchaingo/llms
go get github.com/tmc/langchaingo/llms/openai
go get github.com/joho/godotenv
```

## Configuration

Create a `.env` file in project root:
```
GROQ_API_KEY=your_groq_api_key_here
```

## Project Structure

- Uses OpenAI-compatible Groq API
- Generates poetry about Golang
- Demonstrates streaming text generation
- Supports dynamic generation parameters

## How to Run

```bash
go mod tidy
go run main.go
```

## Key Features

- AI-powered poetry generation
- Real-time streaming output
- Configurable generation parameters
- Llama3 language model
- Env-based configuration management

## Generation Parameters

- Model: `llama3-8b-8192`
- Temperature: 0.8
- Max Tokens: 4096
- Streaming enabled

## Customization Options

- Modify prompt
- Adjust temperature
- Change max tokens
- Experiment with different models

## Streaming Capabilities

- Real-time text generation
- Chunk-by-chunk output
- Simulates conversational AI interaction

## Technologies

- Language: Go
- AI: Groq Cloud / Llama3
- Library: LangChainGo
- Technique: Streaming Generation

## Example Output

```
In the realm of code where logic resides,
Golang stands tall, where clarity abides...
```

## Troubleshooting

- Verify API key
- Check internet connection
- Ensure `.env` file exists
- Update dependencies

## Error Handling

- Comprehensive error logging
- Graceful failure mechanisms
- Environment variable validation