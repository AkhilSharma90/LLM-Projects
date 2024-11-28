# Twilio and Cohere AI-Powered Phone Assistant

This project guides you through creating an AI-powered phone assistant using Twilio for voice call handling and Cohere for natural language processing. The setup allows users to call a phone number, ask questions, and receive AI-generated responses.

## Features
- **Voice Input:** Users interact via voice.
- **AI-Powered Responses:** Cohere AI delivers intelligent, context-aware replies.
- **Twilio Integration:** Manages call routing efficiently.
- **Simple Configuration:** Easily set up using environment variables.

## Prerequisites
- **Go Installation:** Install from [Go's download page](https://go.dev/dl/).
- **Twilio Account:** Register at [Twilio](https://www.twilio.com/).
- **Cohere Account:** Register at [Cohere](https://cohere.ai/).
- **ngrok:** Install from [ngrok's website](https://ngrok.com/) for local server exposure.

## Setup Instructions

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/your-username/your-repo-name.git
   cd your-repo-name
   ```

2. **Install Dependencies:**
   ```bash
   go get github.com/gin-gonic/gin
   go get github.com/joho/godotenv
   go get github.com/twilio/twilio-go
   go get github.com/cohere-ai/cohere-go/v2
   ```

3. **Configure Environment Variables:**
   - Rename `example.env` to `.env`.
   - Edit `.env` to include your credentials and settings:
     ```ini
     TWILIO_ACCOUNT_SID=YOUR_TWILIO_ACCOUNT_SID
     TWILIO_AUTH_TOKEN=YOUR_TWILIO_AUTH_TOKEN
     TWILIO_PHONE_NUMBER=YOUR_TWILIO_PHONE_NUMBER
     COHERE_API_KEY=YOUR_COHERE_API_KEY
     COMPANY_NAME="Your Company Name"  # Example: "Acme Inc."
     SEED_MESSAGE="You are a helpful AI assistant for Your Company Name, answering customer questions. Current time is 2 AM. Here's the question: "
     ```

4. **Run the Application:**
   ```bash
   go run main.go
   ```
   - Starts the server on port 10000.

5. **Expose Local Server with ngrok:**
   ```bash
   ngrok http 10000
   ```
   - Copy the HTTPS forwarding URL provided by ngrok (e.g., `https://xxxxxxx.ngrok.io`).

6. **Configure Twilio Webhook:**
   - Log in to [Twilio](https://www.twilio.com/).
   - Go to your phone numbers; select the one for this project.
   - Under "A Call Comes In," choose "Webhook" and enter the ngrok URL with `/answer` appended (e.g., `https://xxxxxxx.ngrok.io/answer`).
   - Save the configuration.

## Testing

- **Make a Call:** Dial the Twilio phone number you configured.
- **Interact:** You will hear a welcome message; ask a question based on `COMPANY_NAME`.
- **Verification:** Ensure Cohere AI responds to your query. Check:
  - Your ngrok tunnel is active.
  - Twilio webhook is correctly set.
  - Terminal logs for errors.
  - Cohere dashboard for request tracking and usage limits.