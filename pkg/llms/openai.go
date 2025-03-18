package llms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"time"

	"github.com/feiskyer/openai-copilot/pkg/types"
	"github.com/sashabaranov/go-openai"
)

// OpenAIClient is a client for the OpenAI API.
type OpenAIClient struct {
	*openai.Client

	Retries int
	Backoff time.Duration
}

// NewOpenAIClient creates an OpenAI client.
func NewOpenAIClient() (*OpenAIClient, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey != "" {
		config := openai.DefaultConfig(apiKey)
		baseURL := os.Getenv("OPENAI_API_BASE")
		if baseURL != "" {
			config.BaseURL = baseURL
		}

		return &OpenAIClient{
			Retries: 5,
			Backoff: time.Second,
			Client:  openai.NewClientWithConfig(config),
		}, nil
	}

	azureAPIKey := os.Getenv("AZURE_OPENAI_API_KEY")
	azureAPIBase := os.Getenv("AZURE_OPENAI_API_BASE")
	azureAPIVersion := os.Getenv("AZURE_OPENAI_API_VERSION")
	if azureAPIVersion == "" {
		azureAPIVersion = "2025-02-01-preview"
	}
	if azureAPIKey != "" && azureAPIBase != "" {
		config := openai.DefaultConfig(azureAPIKey)
		config.BaseURL = azureAPIBase
		config.APIVersion = azureAPIVersion
		config.APIType = openai.APITypeAzure
		config.AzureModelMapperFunc = func(model string) string {
			return regexp.MustCompile(`[.:]`).ReplaceAllString(model, "")
		}

		return &OpenAIClient{
			Retries: 5,
			Backoff: time.Second,
			Client:  openai.NewClientWithConfig(config),
		}, nil
	}

	return nil, fmt.Errorf("OPENAI_API_KEY or AZURE_OPENAI_API_KEY is not set")
}

// Chat sends a chat completion request to the OpenAI API and returns the response.
func (c *OpenAIClient) Chat(model string, maxTokens int, prompts []openai.ChatCompletionMessage) (*types.ToolPrompt, error) {
	var response types.ToolPrompt

	req := openai.ChatCompletionRequest{
		Model:       model,
		MaxTokens:   maxTokens,
		Temperature: math.SmallestNonzeroFloat32,
		Messages:    prompts,
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
	}
	if model == "o1-mini" || model == "o3-mini" || model == "o1" || model == "o3" {
		req = openai.ChatCompletionRequest{
			Model:               model,
			MaxCompletionTokens: maxTokens,
			Messages:            prompts,
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
		}
	}

	backoff := c.Backoff
	for try := 0; try < c.Retries; try++ {
		resp, err := c.Client.CreateChatCompletion(context.Background(), req)
		if err == nil {
			result := ""
			for _, message := range resp.Choices {
				result += message.Message.Content
			}

			err = json.Unmarshal([]byte(result), &response)
			if err != nil {
				return nil, fmt.Errorf("illegal JSON object: %s", result)
			}

			return &response, nil
		}

		e := &openai.APIError{}

		if errors.As(err, &e) {
			switch e.HTTPStatusCode {
			case 401:
				return nil, err
			case 429, 500:
				time.Sleep(backoff)
				backoff *= 2
				continue
			default:
				return nil, err
			}
		}

		return nil, err
	}

	return nil, fmt.Errorf("OpenAI request throttled after retrying %d times", c.Retries)
}
