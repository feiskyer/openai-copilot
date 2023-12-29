# OpenAI Copilot

Your life Copilot powered by OpenAI (CLI interface for OpenAI with searching).

Features:

* Web access and Google search support without leaving the terminal.
* Automatically execute any steps predicted from prompt instructions.
* Human interactions on uncertain instructions to avoid inappropriate operations.

## Install

Install the copilot with the commands below:

```sh
go install github.com/feiskyer/openai-copilot/cmd/openai-copilot
```

## Setup

* OpenAI API key should be set to `OPENAI_API_KEY` environment variable to enable the ChatGPT feature.
  * `OPENAI_API_BASE` should be set as well for Azure OpenAI service and other self-hosted OpenAI services.
* Google Search API key and CSE ID should be set to `GOOGLE_API_KEY` and `GOOGLE_CSE_ID`.

## How to use

```sh
Usage:
  openai-copilot [flags]

Flags:
  -c, --count-tokens     Print tokens count
  -h, --help             help for openai-copilot
  -t, --max-tokens int   Max tokens for the GPT model (default 1024)
  -m, --model string     OpenAI model to use (default "gpt-4")
  -p, --prompt string    Prompts sent to GPT model for non-interactive mode. If not set, interactive mode is used
  -v, --verbose          Enable verbose output (default true)
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

## Python Version

Please note that the original project (version number < v0.5.0) is written in Python 3 and the codes are in [main](https://github.com/feiskyer/openai-copilot/tree/main) branch.

## Contribution

The project is opensource at github [feiskyer/openai-copilot](https://github.com/feiskyer/openai-copilot) with Apache License.

If you would like to contribute to the project, please follow these guidelines:

1. Fork the repository and clone it to your local machine.
2. Create a new branch for your changes.
3. Make your changes and commit them with a descriptive commit message.
4. Push your changes to your forked repository.
5. Open a pull request to the main repository.
