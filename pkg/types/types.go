package types

// ToolPrompt is the JSON format for the Copilot.
type ToolPrompt struct {
	Question string `json:"question"`
	Thought  string `json:"thought,omitempty"`
	Action   struct {
		Name  string      `json:"name"`
		Input interface{} `json:"input"`
	} `json:"action,omitempty"`
	Observation string `json:"observation,omitempty"`
	FinalAnswer string `json:"final_answer,omitempty"`
}

// Tool is an interface that defines the methods for a tool.
type Tool interface {
	Description() string
	InputSchema() string
	ToolFunc(input string) (string, error)
}
