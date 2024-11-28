package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go/twiml"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
)

func loadEnv() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Load Environment Variables
	loadEnv()

	router := gin.Default()

	router.POST("/answer", incomingCallHandler)
	router.POST("/handle-user-input", handleUserInput)

	router.Run(":10000")
}

// Handles Twilio Webhook for incoming calls
func incomingCallHandler(c *gin.Context) {
	msg := "Welcome! This is a test app, ask me questions related to " + os.Getenv("COMPANY_NAME")

	gather := &twiml.VoiceGather{
		Input:         "speech",
		Language:      "en-US",
		Action:        "/handle-user-input",
		SpeechTimeout: "1",
	}

	say := &twiml.VoiceSay{
		Message: msg,
	}

	twimlResult, err := twiml.Voice([]twiml.Element{say, gather})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Header("Content-Type", "text/xml")
		c.String(http.StatusOK, twimlResult)
	}
}

// Handles Twilio Webhook for user input
func handleUserInput(c *gin.Context) {
	userInput := c.PostForm("SpeechResult")

	client := cohereclient.NewClient(cohereclient.WithToken(os.Getenv("COHERE_API_KEY")))

	response, err := client.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Message: os.Getenv("SEED_MESSAGE") + userInput,
		},
	)

	say := &twiml.VoiceSay{
		Message: response.Text,
	}

	gather := &twiml.VoiceGather{
		Input:         "speech",
		Language:      "en-US",
		Action:        "/handle-user-input",
		SpeechTimeout: "1",
	}

	askMore := &twiml.VoiceSay{
		Message: "Do you have any more questions?",
	}

	twimlResult, err := twiml.Voice([]twiml.Element{say, askMore, gather})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Header("Content-Type", "text/xml")
		c.String(http.StatusOK, twimlResult)
	}
}
