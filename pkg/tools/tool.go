package tools

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/feiskyer/openai-copilot/pkg/types"
	"github.com/mark3labs/mcp-go/client"
)

// PythonREPLTool executes Python code in a REPL environment.
type PythonREPLTool struct{}

// Description returns the description of the tool.
func (p PythonREPLTool) Description() string {
	return "Execute Python code in a REPL environment"
}

// InputSchema returns the input schema for the tool.
func (p PythonREPLTool) InputSchema() string {
	return "Python code in string format to execute"
}

// ToolFunc executes the provided Python code and returns the result.
func (p PythonREPLTool) ToolFunc(script string) (string, error) {
	return PythonREPL(script)
}

// TrivyTool scans container images for vulnerabilities using Trivy.
type TrivyTool struct{}

// Description returns the description of the tool.
func (t TrivyTool) Description() string {
	return "Scan container images for vulnerabilities using Trivy"
}

// InputSchema returns the input schema for the tool.
func (t TrivyTool) InputSchema() string {
	return "Container image in string format to scan"
}

// ToolFunc scans the provided container image and returns the result.
func (t TrivyTool) ToolFunc(image string) (string, error) {
	return Trivy(image)
}

// KubectlTool executes kubectl commands against a Kubernetes cluster.
type KubectlTool struct{}

// Description returns the description of the tool.
func (k KubectlTool) Description() string {
	return "Execute kubectl commands against a Kubernetes cluster"
}

// InputSchema returns the input schema for the tool.
func (k KubectlTool) InputSchema() string {
	return "kubectl command in string format to execute"
}

// ToolFunc executes the provided kubectl command and returns the result.
func (k KubectlTool) ToolFunc(command string) (string, error) {
	return Kubectl(command)
}

// GoogleSearchTool performs web searches using the Google Search API.
type GoogleSearchTool struct{}

// Description returns the description of the tool.
func (g GoogleSearchTool) Description() string {
	return "Search the web using Google"
}

// InputSchema returns the input schema for the tool.
func (g GoogleSearchTool) InputSchema() string {
	return "Search query in string format"
}

// ToolFunc performs a web search using the provided query and returns the result.
func (g GoogleSearchTool) ToolFunc(query string) (string, error) {
	return GoogleSearch(query)
}

// CopilotTools is a map of tool names to tools.
var CopilotTools = map[string]types.Tool{
	"python":  PythonREPLTool{},
	"trivy":   TrivyTool{},
	"kubectl": KubectlTool{},
}

// InitTools initializes the tools.
func InitTools(mcpConfigFile string, verbose bool) (map[string]client.MCPClient, error) {
	if os.Getenv("GOOGLE_API_KEY") != "" && os.Getenv("GOOGLE_CSE_ID") != "" {
		CopilotTools["search"] = GoogleSearchTool{}
	}

	if mcpConfigFile != "" {
		mcpTools, mcpClients, err := GetMCPTools(mcpConfigFile, verbose)
		if err != nil {
			return nil, err
		}

		if verbose {
			tools := ""
			for toolName := range mcpTools {
				tools += fmt.Sprintf("%s, ", toolName)
			}
			color.Green("Enabled MCP tools: %s", strings.TrimRight(tools, ", "))
		}

		for toolName, tool := range mcpTools {
			CopilotTools[toolName] = tool
		}

		return mcpClients, nil
	}

	return nil, nil
}

// GetToolPrompt returns the tool prompt.
func GetToolPrompt() string {
	tools := ""
	for toolName, tool := range CopilotTools {
		tools += fmt.Sprintf("- %s: %s, input schema: %s\n", toolName, tool.Description(), tool.InputSchema())
	}

	return tools
}
