# LangChainGo SQLite Conversation Memory

## Overview

This Go application demonstrates a conversational AI system using:

- OpenAI Language Model
- SQLite for persistent conversation history
- LangChainGo for conversation management

## Critical Prerequisites

### 🚨 IMPORTANT: GCC REQUIREMENT 🚨

**This project REQUIRES GCC to be installed!**

#### Installation by Operating System:

- **macOS**: 
  ```bash
  xcode-select --install
  # or via Homebrew
  brew install gcc
  ```

- **Ubuntu/Debian**:
  ```bash
  sudo apt-get update
  sudo apt-get install build-essential
  ```

- **Windows**:
  - Install MinGW or Windows Subsystem for Linux (WSL)
  - Recommended: [TDM-GCC](http://tdm-gcc.tdragon.net/)

## System Requirements

- Go (version 1.20+)
- GCC Compiler
- OpenAI API Key

## Dependencies

```bash
# SQLite Driver
go get github.com/mattn/go-sqlite3

# LangChainGo Packages
go get github.com/tmc/langchaingo/chains
go get github.com/tmc/langchaingo/llms
go get github.com/tmc/langchaingo/llms/openai
go get github.com/tmc/langchaingo/memory
go get github.com/tmc/langchaingo/memory/sqlite3
```

## Configuration

1. Set OpenAI API Key:
   ```bash
   export OPENAI_API_KEY='your_openai_api_key_here'
   ```

## Project Features

- Persistent conversation history in SQLite
- Conversational context retention
- Stateful interactions with AI
- Dynamic memory management

## How It Works

1. Initializes OpenAI Language Model
2. Creates SQLite database for conversation tracking
3. Stores and retrieves conversation history
4. Allows querying based on conversation context

## Running the Project

```bash
# Ensure you're in the project directory
go mod tidy
go run main.go
```

## Database Management

- Conversation history stored in `history_example.db`
- SQLite provides lightweight, file-based storage
- Conversations persist between application runs

## Example Interaction

The script demonstrates:

- Storing initial conversation context
- Retrieving conversation history
- Answering context-aware questions

## Potential Improvements

- Add encryption for conversation history
- Implement conversation pruning
- Create more advanced memory strategies
- Add error handling for API interactions

## Troubleshooting

### Common Issues

- **Compilation Errors**: Ensure GCC is properly installed
- **API Key Problems**: Verify OpenAI API key is correct
- **SQLite Errors**: Check database permissions

### Debugging

- Use `go run main.go` with verbose error output
- Check system logs
- Verify all dependencies are correctly installed

## Technologies

- Language: Go
- AI: OpenAI
- Database: SQLite
- Library: LangChainGo