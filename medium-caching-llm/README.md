# Ollama LLM Caching Example

## Overview

This Go application demonstrates a simple implementation of LLM (Large Language Model) caching using Ollama's Llama2 model and an in-memory caching mechanism. The project showcases:

- Ollama LLM integration
- In-memory response caching
- Performance optimization through caching

## Features

- Uses Llama2 language model via Ollama
- Implements in-memory caching with 1-minute expiration
- Demonstrates query performance with and without caching
- Displays query results with word wrapping

## Prerequisites

- Go (version 1.20+)
- Ollama installed and configured
- Llama2 model downloaded locally

## Installation

1. **Install Ollama**:

   ```bash
   # Follow Ollama installation instructions for your OS
   # https://ollama.ai/download
   ```

2. **Download Llama2 Model**:

   ```bash
   ollama pull llama2
   ```

3. **Install Go Dependencies**:

   ```bash
   go get github.com/tmc/langchaingo/llms
   go get github.com/tmc/langchaingo/llms/ollama
   go get github.com/tmc/langchaingo/llms/cache
   go get github.com/mitchellh/go-wordwrap
   ```

## Project Structure

- `main.go`: Primary application logic
- Demonstrates LLM query with in-memory caching
- Uses Ollama's Llama2 model
- Implements word-wrapped output

## How It Works

1. Initialize Ollama LLM with Llama2 model
2. Create in-memory cache with 1-minute expiration
3. Wrap LLM with caching mechanism
4. Run multiple queries to demonstrate caching
5. Display query results and execution times

## Running the Project

```bash
go mod tidy
go run main.go
```

## Customization Options

- Change model: Modify `ollama.WithModel()`
- Adjust cache duration: Modify `inmemory.WithExpiration()`
- Alter prompt: Modify the query string
- Adjust temperature: Change `llms.WithTemperature()`

## Example Output

```
## Iteration #0
(First query - fetches from LLM)
Neil Armstrong was the first man to walk on the moon...
(took 2.3s)

## Iteration #1
(Cached result - much faster)
Neil Armstrong was the first man to walk on the moon...
(took 5ms)
```

## Performance Benefits

- First query fetches from LLM (slower)
- Subsequent queries retrieve from cache (faster)
- Reduces unnecessary API/LLM calls
- Improves response time for repeated queries

## Dependencies

- LangChainGo
- Ollama
- Go Wordwrap

## Technologies

- Language: Go
- LLM: Ollama (Llama2)
- Caching: In-memory
- Library: LangChainGo