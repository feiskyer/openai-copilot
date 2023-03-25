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
  --model TEXT       OpenAI model to use for copilot execution, default is
                     gpt-3.5-turbo
  --enable-terminal  Enable Copilot to run programs within terminal. Enable
                     with caution since Copilot may execute inappropriate
                     commands
  --help             Show this message and exit.
```

Here is a conversation sample:

```sh
# openai-copilot
>>> : What is OpenAI?
OpenAI is a research organization consisting of the for-profit corporation OpenAI LP and its parent company, the non-profit OpenAI Inc. The company aims to conduct cutting-edge research and development in artificial intelligence (AI) in a way that benefits humanity as a whole. Some notable achievements of OpenAI include the development of advanced language models such as GPT-3 and the creation of AI agents that are capable of playing complex games such as Dota 2.
>>> : What is its latest product?
OpenAI's latest system is GPT-4, which is OpenAI's most advanced system that produces safer and more useful language models. Additionally, the company also has an API platform that offers their latest models and guides for safety best practices.
>>> :
```

## Contribution

The project is opensource at github [feiskyer/openai-copilot](https://github.com/feiskyer/openai-copilot) with Apache License.

If you would like to contribute to the project, please follow these guidelines:

1. Fork the repository and clone it to your local machine.
2. Create a new branch for your changes.
3. Make your changes and commit them with a descriptive commit message.
4. Push your changes to your forked repository.
5. Open a pull request to the main repository.
