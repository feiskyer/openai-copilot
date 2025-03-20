package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/feiskyer/openai-copilot/pkg/assistants"
	"github.com/feiskyer/openai-copilot/pkg/consts"
	"github.com/feiskyer/openai-copilot/pkg/tools"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

const (
	// VERSION is the version of OpenAI Copilot.
	VERSION = "v0.6.1"
)

var (
	// global flags
	model, prompt string
	maxTokens     int
	maxIterations int
	countTokens   bool
	verbose       bool
	mcpConfigFile string
	version       bool

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "openai-copilot",
		Short: "OpenAI Copilot",
		Run: func(cmd *cobra.Command, args []string) {
			chat()
		},
	}
)

// init initializes the command line flags
func init() {
	rootCmd.PersistentFlags().StringVarP(&model, "model", "m", "gpt-4o", "OpenAI model to use")
	rootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "Prompts sent to GPT model for non-interactive mode. If not set, interactive mode is used")
	rootCmd.PersistentFlags().StringVarP(&mcpConfigFile, "mcp-config", "c", "", "MCP config file")
	rootCmd.PersistentFlags().IntVarP(&maxTokens, "max-tokens", "t", 4000, "Max tokens for the GPT model")
	rootCmd.PersistentFlags().IntVarP(&maxIterations, "max-iterations", "i", 10, "Max iterations for the conversations")
	rootCmd.PersistentFlags().BoolVarP(&countTokens, "count-tokens", "k", false, "Print tokens count")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", true, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "V", false, "Print version information")
}

func chat() {
	var err error
	var response string
	if version {
		fmt.Printf("OpenAI Copilot %s\n", VERSION)
		return
	}

	mcpClients, err := tools.InitTools(mcpConfigFile, verbose)
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer func() {
		for _, mcpClient := range mcpClients {
			mcpClient.Close()
		}
	}()

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: strings.Replace(consts.DefaultPrompt, "{{TOOLS}}", tools.GetToolPrompt(), 1),
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
		if scanner.Err() == io.EOF || scanner.Text() == "exit" || scanner.Text() == "quit" {
			break
		}

		message := scanner.Text()
		if strings.TrimSpace(message) == "" {
			continue
		}

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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
