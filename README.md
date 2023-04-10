# OpenAI Copilot

Your life Copilot powered by OpenAI (CLI interface for OpenAI with searching).

**Status: Experimental**

**Caution: Copilot may generate and execute inappropriate operations, do not use in production environment when enabling terminal execution!**

Features:

* Web access and Google search support without leaving the terminal.
* Automatically execute any steps predicted from prompt instructions.
* Human interactions on uncertain instructions to avoid inappropriate operations.

## Install

Install the copilot with pip command below:

```sh
pip install openai-copilot
```

## Setup

* OpenAI API key should be set to `OPENAI_API_KEY` environment variable to enable the ChatGPT feature.
  * `OPENAI_API_BASE` should be set as well for Azure OpenAI service.
* Google Search API key and CSE ID should be set to `GOOGLE_API_KEY` and `GOOGLE_CSE_ID`.

## How to use

```sh
Usage: openai-copilot [OPTIONS]

  Your life Copilot powered by OpenAI

Options:
  --version          Show the version and exit.
  --verbose          Enable verbose information of copilot execution steps
  --model MODEL      OpenAI model to use for copilot execution, default is
                     gpt-3.5-turbo
  --enable-terminal  Enable Copilot to run programs within terminal. Enable
                     with caution since Copilot may execute inappropriate
                     commands
  --help             Show this message and exit.
```

Here is a conversation sample (user inputs are after `>>>`)):

```sh
# openai-copilot
>>> What is OpenAI?
OpenAI is an artificial intelligence research laboratory consisting of the for-profit corporation OpenAI LP and its parent company, the non-profit OpenAI Inc. The company is dedicated to advancing digital intelligence in a way that is safe and beneficial for humanity as a whole. OpenAI was founded in 2015 by a group of technology leaders including Elon Musk, Sam Altman, Greg Brockman, and Ilya Sutskever. Its mission is to develop and promote friendly AI for the betterment of all humans.
>>> What are the differences between GPT-4 and GPT-3.5?
According to my search results, one of the main differences between GPT-4 and GPT-3.5 is that while GPT-3.5 is a text-to-text model, GPT-4 is more of a data-to-text model. Additionally, GPT-4 has the advantage of providing more creative replies to prompts. However, it's important to note that GPT-4 is not yet released and there is limited information available about it.
>>>
```

## Contribution

The project is opensource at github [feiskyer/openai-copilot](https://github.com/feiskyer/openai-copilot) with Apache License.

If you would like to contribute to the project, please follow these guidelines:

1. Fork the repository and clone it to your local machine.
2. Create a new branch for your changes.
3. Make your changes and commit them with a descriptive commit message.
4. Push your changes to your forked repository.
5. Open a pull request to the main repository.
