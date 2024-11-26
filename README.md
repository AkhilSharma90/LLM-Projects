# Go LLM Tools

This repository contains a collection of Go-based tools and utilities for working with large language models (LLMs). The projects included cover a range of functionality, from interacting with different LLM providers to implementing caching and conversation memory management.

## Project Overview

The repository is structured as follows:

- **anthropic-completion**: This project demonstrates how to use the Anthropic LLM service to generate text completions with streaming output.
- **anthropic-tool-call**: This project shows how to use the Anthropic LLM service in a tool-based interaction, where the model can request external functions to be executed.
- **caching-llm**: This project illustrates how to implement caching for an LLM, using an in-memory cache backend to improve performance for repeated queries.
- **chain-conversation-memory**: This project demonstrates the use of a conversation chain and SQLite-based chat message history to maintain context across multiple interactions.
- **cohere-completion**: This project shows how to use the Cohere LLM service to generate text completions.
- **google-completion**: This project shows how to use the Google(gemini) LLM service to generate text completions.
- **google-streaming**: This project demonstrates the use of the Google(gemini) LLM service with streaming output.
- **groq-completion**: This project demonstrates the use of the GROQ LLM service, which is based on the OpenAI API but with a different authentication mechanism.
- **mistral-completion**: This project shows how to use the Mistral LLM service to generate text completions, with support for both streaming and non-streaming output.
- **ollama-streaming**: This project demonstrates the use of the Ollama LLM service (which is based on the LLaMA model) with streaming output.
- **openai-completion**: This project shows how to use the OpenAI LLM service to generate text completions, with the ability to control parameters like temperature and stop words.
- **openai-o1**: This project demonstrates the use of the OpenAI O1 model, which is a specialized model optimized for certain use cases.

Each project has a `main.go` file that serves as the entry point and contains the core functionality.

## Getting Started

To run the projects in this repository, you'll need to have the following prerequisites:

- Go 1.18 or later installed on your system.
- API keys or credentials for the LLM services you plan to use (e.g., Anthropic, Cohere, GROQ, Mistral, OpenAI).

Once you have the necessary prerequisites, you can clone the repository and navigate to the desired project directory. Each project's `main.go` file can be run using the `go run` command.

For example, to run the `anthropic-completion` project:

```bash
cd anthropic-completion
go run main.go
```