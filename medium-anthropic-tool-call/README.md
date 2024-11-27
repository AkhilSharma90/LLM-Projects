# Weather AI Assistant with LangChainGo and Anthropic Toolcalls

## Overview

This is a sophisticated Go application that demonstrates an AI-powered weather assistant using:
- Anthropic Claude AI (Claude-3-Haiku)
- Open-Meteo Weather API
- LangChainGo for tool-based interactions
- Conversational AI with function calling

## Features

- Fetch real-time weather for specific cities (Pune, Mumbai)
- AI-powered conversational weather queries
- Dynamic tool-based function calling
- Supports multiple interaction scenarios

## Prerequisites

- Go (version 1.20+)
- Anthropic API Key

## Dependencies

```bash
go get github.com/tmc/langchaingo/llms
go get github.com/tmc/langchaingo/llms/anthropic
```

## Environment Setup

Set your Anthropic API key:
```bash
export ANTHROPIC_API_KEY='your_anthropic_api_key_here'
```

## Project Structure

- `main.go`: Primary application logic
- Implements weather fetching via Open-Meteo API
- Uses Claude AI for intelligent interactions
- Supports conversational weather queries

## Workflow

1. User asks about weather in a city
2. AI determines appropriate function to call
3. Fetches real-time weather data
4. Provides contextual weather information
5. Supports follow-up questions and comparisons

## Supported Locations

Currently hardcoded to support:
- Pune
- Mumbai

## How to Run

```bash
go mod tidy
go run main.go
```

## Example Interactions

The script demonstrates:
- Querying weather for Pune
- Querying weather for Mumbai
- Comparing weather conditions
- AI-driven conversational flow

## Customization

- Add more cities to `getCityCoordinates()`
- Expand `weatherCodeMap` for more detailed descriptions
- Modify tools and interaction patterns

## Error Handling

- Validates API key
- Handles network and parsing errors
- Provides fallback for unknown locations

## Technologies

- Language: Go
- AI: Anthropic Claude-3-Haiku
- Weather API: Open-Meteo
- Library: LangChainGo
