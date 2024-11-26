package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

// WeatherResponse represents the structure of weather data from Open-Meteo API
type WeatherResponse struct {
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
	Timezone       string         `json:"timezone"`
	CurrentWeather CurrentWeather `json:"current_weather"`
}

type CurrentWeather struct {
	Temperature   float64 `json:"temperature"`
	Windspeed     float64 `json:"windspeed"`
	Winddirection float64 `json:"winddirection"`
	WeatherCode   int     `json:"weathercode"`
	Time          string  `json:"time"`
}

// getCityCoordinates fetches coordinates for a given city
func getCityCoordinates(city string) (float64, float64, error) {
	// Hardcoded coordinates for Pune and Mumbai
	cityCoords := map[string]struct{ lat, lon float64 }{
		"pune":   {18.5204, 73.8567},
		"mumbai": {19.0760, 72.8777},
	}

	// Normalize city name (lowercase)
	normalizedCity := strings.ToLower(city)

	// Look up coordinates
	coords, exists := cityCoords[normalizedCity]
	if !exists {
		return 0, 0, fmt.Errorf("coordinates not found for city: %s", city)
	}

	return coords.lat, coords.lon, nil
}

// getCurrentWeather fetches real-time weather data for a given city
func getCurrentWeather(location string, unit string) (string, error) {
	// Get coordinates for the city
	lat, lon, err := getCityCoordinates(location)
	if err != nil {
		return "", err
	}

	// Construct URL for Open-Meteo API
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&current_weather=true", lat, lon)

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	// Check for API errors
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %s", string(body))
	}

	// Parse JSON response
	var weatherData WeatherResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}

	// Map weather code to description
	weatherCodeMap := map[int]string{
		0:  "Clear sky",
		1:  "Mainly clear",
		2:  "Partly cloudy",
		3:  "Overcast",
		45: "Foggy",
		48: "Depositing rime fog",
		51: "Light drizzle",
		53: "Moderate drizzle",
		55: "Dense drizzle",
	}

	// Get weather description
	weatherDesc := "Unknown conditions"
	if desc, exists := weatherCodeMap[weatherData.CurrentWeather.WeatherCode]; exists {
		weatherDesc = desc
	}

	// Format weather information
	weatherInfo := fmt.Sprintf("%.1f°C and %s (Wind: %.1f km/h)",
		weatherData.CurrentWeather.Temperature,
		weatherDesc,
		weatherData.CurrentWeather.Windspeed)

	// Marshal to JSON for consistency with previous implementation
	b, err := json.Marshal(weatherInfo)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {
	llm, err := anthropic.New(anthropic.WithModel("claude-3-haiku-20240307"))
	if err != nil {
		log.Fatal(err)
	}

	// Sending initial message to the model, with a list of available tools.
	ctx := context.Background()
	messageHistory := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "What is the weather like in Pune?"),
	}

	fmt.Println("Querying for weather in Pune..")
	resp, err := llm.GenerateContent(ctx, messageHistory, llms.WithTools(availableTools))
	if err != nil {
		log.Fatal(err)
	}

	// Execute tool calls requested by the model
	messageHistory = executeToolCalls(ctx, llm, messageHistory, resp)

	// Send query to the model again, this time with a history containing its
	// request to invoke a tool and our response to the tool call.
	fmt.Println("Querying with tool response...")
	resp, err = llm.GenerateContent(ctx, messageHistory, llms.WithTools(availableTools))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Choices[0].Content)

	// populate ai response
	assistantResponse := llms.TextParts(llms.ChatMessageTypeAI, resp.Choices[0].Content)
	messageHistory = append(messageHistory, assistantResponse)

	fmt.Println("asking about Mumbai...")
	// Human asks about Mumbai
	humanQuestion := llms.TextParts(llms.ChatMessageTypeHuman, "How about the weather in Mumbai?")
	messageHistory = append(messageHistory, humanQuestion)

	// Send Request
	resp, err = llm.GenerateContent(ctx, messageHistory, llms.WithTools(availableTools))
	if err != nil {
		log.Fatal(err)
	}

	// Perform Tool call
	messageHistory = executeToolCalls(ctx, llm, messageHistory, resp)
	fmt.Println("Querying with tool response...")
	resp, err = llm.GenerateContent(ctx, messageHistory, llms.WithTools(availableTools))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Choices[0].Content)

	// populate ai response
	assistantResponse = llms.TextParts(llms.ChatMessageTypeAI, resp.Choices[0].Content)
	messageHistory = append(messageHistory, assistantResponse)

	// Compare responses
	humanQuestion = llms.TextParts(llms.ChatMessageTypeHuman, "How do these compare?")
	messageHistory = append(messageHistory, humanQuestion)

	// Send Request
	resp, err = llm.GenerateContent(ctx, messageHistory, llms.WithTools(availableTools))
	if err != nil {
		log.Fatal(err)
	}
	// Perform Tool call
	messageHistory = executeToolCalls(ctx, llm, messageHistory, resp)
	fmt.Println("Asking for comparison...")
	resp, err = llm.GenerateContent(ctx, messageHistory, llms.WithTools(availableTools))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Choices[0].Content)
}

// executeToolCalls remains the same as in the previous example
func executeToolCalls(ctx context.Context, llm llms.Model, messageHistory []llms.MessageContent, resp *llms.ContentResponse) []llms.MessageContent {
	for _, choice := range resp.Choices {
		for _, toolCall := range choice.ToolCalls {

			// Append tool_use to messageHistory
			assistantResponse := llms.MessageContent{
				Role: llms.ChatMessageTypeAI,
				Parts: []llms.ContentPart{
					llms.ToolCall{
						ID:   toolCall.ID,
						Type: toolCall.Type,
						FunctionCall: &llms.FunctionCall{
							Name:      toolCall.FunctionCall.Name,
							Arguments: toolCall.FunctionCall.Arguments,
						},
					},
				},
			}
			messageHistory = append(messageHistory, assistantResponse)

			switch toolCall.FunctionCall.Name {
			case "getCurrentWeather":
				var args struct {
					Location string `json:"location"`
					Unit     string `json:"unit"`
				}
				if err := json.Unmarshal([]byte(toolCall.FunctionCall.Arguments), &args); err != nil {
					log.Fatal(err)
				}

				// Perform Function Calling
				response, err := getCurrentWeather(args.Location, args.Unit)
				if err != nil {
					log.Fatal(err)
				}

				// Append tool_result to messageHistory
				weatherCallResponse := llms.MessageContent{
					Role: llms.ChatMessageTypeTool,
					Parts: []llms.ContentPart{
						llms.ToolCallResponse{
							ToolCallID: toolCall.ID,
							Name:       toolCall.FunctionCall.Name,
							Content:    response,
						},
					},
				}
				messageHistory = append(messageHistory, weatherCallResponse)
			default:
				log.Fatalf("Unsupported tool: %s", toolCall.FunctionCall.Name)
			}
		}
	}

	return messageHistory
}

var availableTools = []llms.Tool{
	{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "getCurrentWeather",
			Description: "Get the current weather in a given location",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"location": map[string]any{
						"type":        "string",
						"description": "The city and state, e.g. San Francisco, CA",
					},
					"unit": map[string]any{
						"type": "string",
						"enum": []string{"fahrenheit", "celsius"},
					},
				},
				"required": []string{"location"},
			},
		},
	},
}
