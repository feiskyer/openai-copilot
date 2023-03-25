#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import click
from openai_copilot.llm import init_openai
from openai_copilot.agent import CopilotLLM


@click.command()
@click.version_option()
@click.option("--verbose", is_flag=True, default=False, help="Enable verbose information of copilot execution steps")
@click.option("--model", default="gpt-3.5-turbo", help="OpenAI model to use for copilot execution, default is gpt-3.5-turbo")
@click.option("--enable-terminal", is_flag=True, default=False, help="Enable Copilot to run programs within terminal. Enable with caution since Copilot may execute inappropriate commands")
def cli(verbose, model, enable_terminal):
    '''Your life Copilot powered by OpenAI'''
    init_openai()
    chain = CopilotLLM(
        verbose=verbose, model=model, enable_terminal=enable_terminal)
    while True:
        instructions = click.prompt(">>> ")
        result = chain.run(instructions)
        click.echo(click.style(result, fg='green'))


def main():
    cli()


if __name__ == "__main__":
    main()
