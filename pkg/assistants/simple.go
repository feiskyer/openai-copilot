package assistants

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/feiskyer/openai-copilot/pkg/llms"
	"github.com/feiskyer/openai-copilot/pkg/tools"
	"github.com/sashabaranov/go-openai"
)

const (
	defaultMaxIterations = 10
)

// Assistant is the simplest AI assistant.
func Assistant(model string, prompts []openai.ChatCompletionMessage, maxTokens int, countTokens bool, verbose bool, maxIterations int) (result string, chatHistory []openai.ChatCompletionMessage, err error) {
	chatHistory = prompts
	if len(chatHistory) == 0 {
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

	if verbose {
		color.Blue("Chatting with LLM\n")
	}

	resp, err := client.Chat(model, maxTokens, chatHistory)
	if err != nil {
		return "", chatHistory, fmt.Errorf("chat completion error: %v", err)
	}

	response, err := json.Marshal(resp)
	if err != nil {
		return "", chatHistory, fmt.Errorf("failed to marshal response: %v", err)
	}
	chatHistory = append(chatHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: string(response),
	})

	if verbose {
		color.Cyan("Initial response from LLM:\n%s\n\n", resp)
	}

	iterations := 0
	if maxIterations <= 0 {
		maxIterations = defaultMaxIterations
	}
	for {
		iterations++

		if verbose {
			color.Cyan("Thought: %s\n\n", resp.Thought)
		}

		if iterations > maxIterations {
			color.Red("Max iterations reached")
			return resp.FinalAnswer, chatHistory, nil
		}

		if resp.FinalAnswer != "" {
			if verbose {
				color.Cyan("Final answer: %s\n\n", resp.FinalAnswer)
			}
			return resp.FinalAnswer, chatHistory, nil
		}

		if resp.Action.Name != "" {
			input, ok := resp.Action.Input.(string)
			if !ok {
				inputBytes, err := json.Marshal(resp.Action.Input)
				if err != nil {
					return "", chatHistory, fmt.Errorf("failed to marshal tool input: %v", err)
				}
				input = string(inputBytes)
			}
			if verbose {
				color.Blue("Executing tool %s\n", resp.Action.Name)
				color.Cyan("Invoking %s tool with params: \n============\n%s\n============\n\n", resp.Action.Name, input)
			}
			ret, err := tools.CopilotTools[resp.Action.Name].ToolFunc(input)
			observation := strings.TrimSpace(ret)
			if err != nil {
				observation = fmt.Sprintf("Tool %s failed with ret %s and error %s. Considering refine the inputs for the tool.", resp.Action.Name, ret, err)
			}
			if verbose {
				color.Cyan("Observation: %s\n\n", observation)
			}

			// Constrict the prompt to the max tokens allowed by the model.
			// This is required because the tool may have generated a long output.
			observation = llms.ConstrictPrompt(observation, model)
			resp.Observation = observation
			assistantMessage, _ := json.Marshal(resp)
			chatHistory = append(chatHistory, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: string(assistantMessage),
			})
			// Constrict the chat history to the max tokens allowed by the model.
			// This is required because the chat history may have grown too large.
			chatHistory = llms.ConstrictMessages(chatHistory, model)

			// Start next iteration of LLM chat.
			if verbose {
				color.Blue("Chatting with LLM\n")
			}

			resp, err = client.Chat(model, maxTokens, chatHistory)
			if err != nil {
				return "", chatHistory, fmt.Errorf("chat completion error: %v", err)
			}

			response, err := json.Marshal(resp)
			if err != nil {
				return "", chatHistory, fmt.Errorf("failed to marshal response: %v", err)
			}
			chatHistory = append(chatHistory, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: string(response),
			})
			if verbose {
				color.Cyan("Intermediate response from LLM: %s\n\n", resp)
			}
		}
	}
}
