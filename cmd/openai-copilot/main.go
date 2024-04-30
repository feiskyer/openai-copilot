package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/feiskyer/openai-copilot/pkg/assistants"
	"github.com/feiskyer/openai-copilot/pkg/consts"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var (
	// global flags
	model, prompt string
	maxTokens     int
	maxIterations int
	countTokens   bool
	verbose       bool

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "openai-copilot",
		Short: "OpenAI Copilot",
		Run: func(cmd *cobra.Command, args []string) {
			chat()
		},
	}
)

func chat() {
	var err error
	var response string
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: consts.DefaultPrompt,
		},
	}

	// Non-interactive mode
	if prompt != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		})
		response, _, err = assistants.Assistant(model, messages, maxTokens, countTokens, verbose, maxIterations)
		if err != nil {
			color.Red(err.Error())
			return
		}

		fmt.Printf("%s\n\n", response)
		return
	}

	// Interactive mode
	color.New(color.FgYellow).Printf("You: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: message,
		})
		response, messages, err = assistants.Assistant(model, messages, maxTokens, countTokens, verbose, maxIterations)
		if err != nil {
			color.Red(err.Error())
			continue
		}

		color.New(color.FgYellow).Printf("AI: ")
		fmt.Printf("%s\n\n", response)
		color.New(color.FgYellow).Printf("You: ")
	}
}

// init initializes the command line flags
func init() {
	rootCmd.PersistentFlags().StringVarP(&model, "model", "m", "gpt-4", "OpenAI model to use")
	rootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "Prompts sent to GPT model for non-interactive mode. If not set, interactive mode is used")
	rootCmd.PersistentFlags().IntVarP(&maxTokens, "max-tokens", "t", 1024, "Max tokens for the GPT model")
	rootCmd.PersistentFlags().IntVarP(&maxIterations, "max-iterations", "i", 3, "Max iterations for the conversations")
	rootCmd.PersistentFlags().BoolVarP(&countTokens, "count-tokens", "c", false, "Print tokens count")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", true, "Enable verbose output")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
