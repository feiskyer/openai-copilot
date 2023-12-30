package assistants

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/fatih/color"
	"github.com/feiskyer/openai-copilot/pkg/llms"
	"github.com/feiskyer/openai-copilot/pkg/tools"
	"github.com/sashabaranov/go-openai"
)

// Assistant is the simplest AI assistant.
func Assistant(model string, prompts []openai.ChatCompletionMessage, maxTokens int, countTokens bool, verbose bool) (result string, chatHistory []openai.ChatCompletionMessage, err error) {
	chatHistory = prompts
	if len(prompts) == 0 {
		return "", nil, fmt.Errorf("prompts cannot be empty")
	}

	client, err := llms.NewOpenAIClient()
	if err != nil {
		return "", nil, fmt.Errorf("unable to get OpenAI client: %v", err)
	}

	defer func() {
		if countTokens {
			count := llms.NumTokensFromMessages(chatHistory, model)
			color.Green("Total tokens: %d\n\n", count)
		}
	}()

	req := openai.ChatCompletionRequest{
		Model:       model,
		MaxTokens:   maxTokens,
		Temperature: math.SmallestNonzeroFloat32,
		Messages:    chatHistory,
	}
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", chatHistory, fmt.Errorf("chat completion error: %v", err)
	}
	chatHistory = append(chatHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: string(resp.Choices[0].Message.Content),
	})

	if verbose {
		color.Cyan("Initial response from LLM:\n%s\n\n", resp.Choices[0].Message.Content)
	}

	var toolPrompt tools.ToolPrompt
	if err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &toolPrompt); err != nil {
		if verbose {
			color.Cyan("Unable to parse tool from prompts (%s), assuming got final answer", resp.Choices[0].Message.Content)
		}
		return resp.Choices[0].Message.Content, chatHistory, nil
	}

	iterations := 0
	maxIterations := 10
	for {
		if verbose {
			color.Cyan("Thought: %s\n\n", toolPrompt.Thought)
		}

		if iterations > maxIterations {
			color.Red("Max iterations reached")
			break
		}
		iterations++

		if toolPrompt.FinalAnswer != "" {
			if verbose {
				color.Cyan("Final answer: %s\n\n", toolPrompt.FinalAnswer)
			}
			return toolPrompt.FinalAnswer, chatHistory, nil
		}

		if toolPrompt.Action.Name != "" {
			if verbose {
				color.Cyan("Invoking %s tool with inputs: \n============\n%s\n============\n\n", toolPrompt.Action.Name, toolPrompt.Action.Input)
			}
			ret, err := tools.CopilotTools[toolPrompt.Action.Name](toolPrompt.Action.Input)
			if err != nil {
				return "", chatHistory, fmt.Errorf("tool %s error: %v", toolPrompt.Action.Name, err)
			}

			observation := strings.TrimSpace(ret)
			if verbose {
				color.Cyan("Observation: %s\n\n", observation)
			}

			// Constrict the prompt to the max tokens allowed by the model.
			// This is required because the tool may have generated a long output.
			observation = llms.ConstrictPrompt(observation, model, maxTokens)

			toolPrompt.Observation = observation
			assistantMessage, _ := json.Marshal(toolPrompt)
			chatHistory = append(chatHistory, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: string(assistantMessage),
			})
			req := openai.ChatCompletionRequest{
				Model:       openai.GPT4,
				MaxTokens:   maxTokens,
				Temperature: math.SmallestNonzeroFloat32,
				Messages:    chatHistory,
			}
			resp, err = client.CreateChatCompletion(context.Background(), req)
			if err != nil {
				return "", chatHistory, fmt.Errorf("chat completion error: %v", err)
			}
			chatHistory = append(chatHistory, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: string(resp.Choices[0].Message.Content),
			})

			// Constrict the chat history to the max tokens allowed by the model.
			// This is required because the chat history may have grown too large.
			chatHistory = llms.ConstrictMessages(chatHistory, model, maxTokens)

			if err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &toolPrompt); err != nil {
				if verbose {
					color.Cyan("Unable to parse tool from prompts (%s), assuming got final answer", resp.Choices[0].Message.Content)
				}
				return resp.Choices[0].Message.Content, chatHistory, nil
			}
		}
	}

	return "", chatHistory, nil
}
