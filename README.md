# OpenAI Copilot

Your life Copilot powered by OpenAI (CLI interface for OpenAI with searching).

Features:

* Web access and Google search support without leaving the terminal.
* Automatically execute any steps predicted from prompt instructions.
* Human interactions on uncertain instructions to avoid inappropriate operations.

## Install

Install the copilot with the commands below:

```sh
pip install openai-copilot
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
  -h, --help             help for openai-copilot
  -t, --max-tokens int   Max tokens for the GPT model (default 1024)
  -m, --model string     OpenAI model to use (default "gpt-4")
  -p, --prompt string    Prompts sent to GPT model
```

### Non-interactive mode

```sh
$ openai-copilot -p 'What is OpenAI?'
Thought: OpenAI is an artificial intelligence research lab made up of both for-profit and non-profit arms. I can provide a more detailed explanation.

OpenAI is an artificial intelligence research lab made up of the for-profit OpenAI LP and its parent company, the non-profit OpenAI Inc. Their mission is to ensure that artificial general intelligence (AGI) benefits all of humanity. They aim to directly build safe and beneficial AGI, but are also committed to aiding others in achieving this outcome.
```

### Interactive mode

Here is a conversation sample (user inputs are after `>>>`)):

```sh
# openai-copilot
>>> What is OpenAI?
OpenAI is an artificial intelligence research laboratory consisting of the for-profit corporation OpenAI LP and its parent company, the non-profit OpenAI Inc. The company is dedicated to advancing digital intelligence in a way that is safe and beneficial for humanity as a whole. OpenAI was founded in 2015 by a group of technology leaders including Elon Musk, Sam Altman, Greg Brockman, and Ilya Sutskever. Its mission is to develop and promote friendly AI for the betterment of all humans.
>>> What are the differences between GPT-4 and GPT-3.5?
According to my search results, one of the main differences between GPT-4 and GPT-3.5 is that while GPT-3.5 is a text-to-text model, GPT-4 is more of a data-to-text model. Additionally, GPT-4 has the advantage of providing more creative replies to prompts. However, it's important to note that GPT-4 is not yet released and there is limited information available about it.
>>>
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
