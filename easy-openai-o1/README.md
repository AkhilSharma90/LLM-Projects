# LangChainGo OpenAI Content Generation CLI

## Overview

This Go application demonstrates content generation using the LangChainGo library with OpenAI's language models. The CLI tool allows you to generate content by specifying different OpenAI models and provides flexible configuration options.

## Features

- Command-line model selection
- Integration with OpenAI's language models
- Flexible content generation
- JSON output of generation information

## Prerequisites

- Go (version 1.16 or later)
- OpenAI API Key

## Installation

#### Clone the repository:

```bash
git clone <your-repository-url>
cd <project-directory>
```

#### Set up your OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY='your-api-key-here'
```

## Dependencies

This project uses the following key dependencies:

- `github.com/tmc/langchaingo/llms`
- `github.com/tmc/langchaingo/llms/openai`

#### Dependency Installation

Install dependencies using:

```bash
go mod init your-project-name
go get github.com/tmc/langchaingo/llms
go get github.com/tmc/langchaingo/llms/openai
```

## Usage

Run the application with optional model selection:

```bash
# Use default model (o1-preview)
go run main.go

# Specify a different model
go run main.go --model o1-mini
```

### Available Flags

- `--model`: Specify the OpenAI model to use  
  Default: `o1-preview`  
  Examples: `o1-preview`, `o1-mini`

### Example Output

The application will:

- Generate content based on a predefined prompt
- Print the generated content
- Output generation metadata in JSON format

## Configuration

You can customize the generation by modifying:

- Model selection via command-line flag
- Prompt content in the source code
- Generation parameters like temperature

## Error Handling

The application includes basic error handling:

- Checks for errors when initializing the OpenAI client
- Handles potential errors during content generation