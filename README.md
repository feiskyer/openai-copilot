# OpenAI Copilot

Your life Copilot powered by LLM models (CLI interface for LLM models with MCP tools).

Features:

* OpenAI, Azure OpenAI, Anthropic Claude, Google Gemini and OpenAI-compatible services support.
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
| AZURE_OPENAI_API_VERSION | Azure OpenAI API version. Default is 2025-03-01-preview |

### LLM Integrations

<details>
<summary>OpenAI</summary>

Set the OpenAI [API key](https://platform.openai.com/account/api-keys) as the `OPENAI_API_KEY` environment variable to enable OpenAI functionality.
</details>

<details>
<summary>Anthropic Claude</summary>

Anthropic Claude provides an [OpenAI compatible API](https://docs.anthropic.com/en/api/openai-sdk), so it could be used by using following config:

- `OPENAI_API_KEY=<your-anthropic-key>`
- `OPENAI_API_BASE='https://api.anthropic.com/v1/'`

</details>

<details>

<summary>Azure OpenAI</summary>

For [Azure OpenAI service](https://learn.microsoft.com/en-us/azure/cognitive-services/openai/quickstart?tabs=command-line&pivots=rest-api#retrieve-key-and-endpoint), set the following environment variables:

- `AZURE_OPENAI_API_KEY=<your-api-key>`
- `AZURE_OPENAI_API_BASE=https://<replace-this>.openai.azure.com/`
- `AZURE_OPENAI_API_VERSION=2025-03-01-preview`

</details>

<details>
<summary>Google Gemini</summary>

Google Gemini provides an OpenAI compatible API, so it could be used by using following config:

- `OPENAI_API_KEY=<your-google-ai-key>`
- `OPENAI_API_BASE='https://generativelanguage.googleapis.com/v1beta/openai/'`

</details>

<details>
<summary>Ollama or other OpenAI compatible LLMs</summary>

For Ollama or other OpenAI compatible LLMs, set the following environment variables:

- `OPENAI_API_KEY=<your-api-key>`
- `OPENAI_API_BASE='http://localhost:11434/v1'` (or your own base URL)

</details>

## How to use

Setup the OpenAI or AzureOpenAI environment and then run the following openai-copilot command:

```sh
# You can optionally config a different model and
# enable MCP integration (MCP format is same as Claude Desktop)
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

### Use MCP

```sh
# Use MCP servers configured from Claude Desktop
openai-copilot -c ~/Library/Application\ Support/Claude/claude_desktop_config.json
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
