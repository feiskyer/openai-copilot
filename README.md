# OpenAI Copilot

Your life Copilot powered by LLM models (CLI interface for LLM models with MCP tools).

Features:

* OpenAI, Azure OpenAI, and OpenAI-compatible services support.
* Web access and Google search support without leaving the terminal.
* Automatically execute any steps predicted from prompt instructions.
* Model Context Protocol (MCP) support for bridging external tools.

## Install

Install the copilot with the commands below:

```sh
go install github.com/feiskyer/openai-copilot/cmd/openai-copilot@latest
```

## Setup

| Environment | Description |
|-------------|-------------|
| OPENAI_API_KEY | OpenAI key. Required for OpenAI or OpenAI-compatible services. |
| OPENAI_API_BASE | OpenAI Base URL. Optional. |
| AZURE_OPENAI_API_KEY | Azure OpenAI API key. Required for Azure OpenAI service. |
| AZURE_OPENAI_API_BASE | Azure OpenAI Base URL. Required for Azure OpenAI service. |
| GOOGLE_API_KEY | Google API key. Required for Google search. |
| GOOGLE_CSE_ID | Google Custom Search Engine ID. Required for Google search. |

## How to use

Setup the OpenAI or AzureOpenAI environment and then run the following openai-copilot command:

```sh
# You can optionally config a different model and 
# enable MCP integration (MCP format is same as Claude app)
openai-copilot [-m gpt-4o] [-c <mcp-config-file>]
```

Here is the full usage of the openai-copilot command:

```sh
OpenAI Copilot

Usage:
  openai-copilot [flags]

Flags:
  -k, --count-tokens         Print tokens count
  -h, --help                 help for openai-copilot
  -i, --max-iterations int   Max iterations for the conversations (default 10)
  -t, --max-tokens int       Max tokens for the GPT model (default 4000)
  -c, --mcp-config string    MCP config file
  -m, --model string         OpenAI model to use (default "gpt-4o")
  -p, --prompt string        Prompts sent to GPT model for non-interactive mode. If not set, interactive mode is used
  -v, --verbose              Enable verbose output (default true)
```

### Interactive mode

Here is a conversation sample (user inputs are after `You:`)):

```sh
$ openai-copilot --verbose=false
You: What is OpenAI?
AI: OpenAI is an artificial intelligence research lab, which includes a for-profit arm, OpenAI LP, and its parent company, the non-profit OpenAI Inc. Their mission is to ensure that artificial general intelligence (AGI) benefits all of humanity. They aim to build safe and beneficial AGI, and are also committed to aiding others in achieving this outcome.

You:
```

### Non-interactive mode

```sh
$ openai-copilot -p 'What is OpenAI?'
Initial response from LLM:
{
 "question": "What is OpenAI?",
 "thought": "OpenAI is a well-known organization in the field of artificial intelligence. I should provide a brief description of it.",
 "action": {
  "name": "search",
  "input": "OpenAI"
 },
 "observation": "OpenAI is an artificial intelligence research lab consisting of the for-profit arm OpenAI LP and its parent company, the non-profit OpenAI Inc. OpenAI's mission is to ensure that artificial general intelligence (AGI) benefits all of humanity. They aim to build safe and beneficial AGI directly, but are also committed to aiding others in achieving this outcome.",
 "final_answer": "OpenAI is an artificial intelligence research lab made up of a for-profit arm, OpenAI LP, and its parent company, the non-profit OpenAI Inc. Their mission is to ensure that artificial general intelligence (AGI) benefits all of humanity. They aim to directly build safe and beneficial AGI, but are also committed to aiding others in achieving this outcome."
}

Thought: OpenAI is a well-known organization in the field of artificial intelligence. I should provide a brief description of it.

Final answer: OpenAI is an artificial intelligence research lab made up of a for-profit arm, OpenAI LP, and its parent company, the non-profit OpenAI Inc. Their mission is to ensure that artificial general intelligence (AGI) benefits all of humanity. They aim to directly build safe and beneficial AGI, but are also committed to aiding others in achieving this outcome.

AI: OpenAI is an artificial intelligence research lab made up of a for-profit arm, OpenAI LP, and its parent company, the non-profit OpenAI Inc. Their mission is to ensure that artificial general intelligence (AGI) benefits all of humanity. They aim to directly build safe and beneficial AGI, but are also committed to aiding others in achieving this outcome.
```

## Contribution

The project is opensource at github [feiskyer/openai-copilot](https://github.com/feiskyer/openai-copilot) with Apache License.

If you would like to contribute to the project, please follow these guidelines:

1. Fork the repository and clone it to your local machine.
2. Create a new branch for your changes.
3. Make your changes and commit them with a descriptive commit message.
4. Push your changes to your forked repository.
5. Open a pull request to the main repository.
