# Gemini-Powered Chatbot Recruiter

This project demonstrates a simple chat application that simulates a recruiter for a specific candidate, Mark Eschbach, using Google's Gemini Pro model. The chatbot utilizes a knowledge base to answer questions about Mark's experience and qualifications effectively.

## Prerequisites

To get started, you'll need to make sure the following are set up:

1. **Go:** Ensure Go is installed and configured on your system.

2. **Google Cloud Project and Gemini API Key:**
   - Create a Google Cloud Project.
   - Enable the Gemini API.
   - Obtain an API key for authentication.

3. **Go Packages:**

   Install the necessary Go packages by running:

   ```bash
   go get github.com/gorilla/websocket
   go get github.com/google/generative-ai-go/genai
   go get google.golang.org/api/option
   ```

4. **API Key Setup:**

   Set the `GEMINI_API_KEY` environment variable to your Gemini API key:

   ```bash
   export GEMINI_API_KEY="YOUR_ACTUAL_API_KEY"
   ```

   On Windows, use:

   ```bash
   set GEMINI_API_KEY="YOUR_ACTUAL_API_KEY"
   ```

5. **Knowledge Base:**

   The file `data.txt` contains Mark's resume information and serves as the knowledge base for the chatbot. You can customize this file with different information as needed.

## Running the Project

1. Navigate to the project directory.

2. Start the server by executing:

   ```bash
   go run .
   ```

3. Open a web browser and visit `http://localhost:8080`.

## Usage

Type your questions about Mark Eschbach into the text box at the bottom of the page and click "Send." The chatbot will respond based on the information available in `data.txt`.

## How It Works

The project comprises several key components:

- **main.go:** Initializes the server, WebSocket connections, and the Gemini API client.
- **home.html:** Provides the client-side chat interface.
- **client.go:** Manages individual WebSocket connections.
- **hub.go:** Acts as a central message hub, broadcasting messages between clients.
- **agent.go:** Interacts with the Gemini API, using the contents of `data.txt` as context.
- **data.txt:** The knowledge base containing Mark Eschbach's information.

## Key Concepts

- **WebSockets:** Used for real-time communication between the client and server.
- **Gemini Pro:** A powerful language model by Google that enables the chatbot's conversational abilities.
- **Contextual Chat:** The chatbot uses the provided resume information as context to answer questions relevantly.