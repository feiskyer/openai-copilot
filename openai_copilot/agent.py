# -*- coding: utf-8 -*-
import os

from langchain.agents import AgentType, Tool, initialize_agent, load_tools
from langchain.chat_models import ChatOpenAI
from langchain.memory import ConversationSummaryBufferMemory
from langchain.tools.python.tool import PythonREPLTool
from langchain.utilities import GoogleSearchAPIWrapper
from openai_copilot.output import ChatOutputParser


class CopilotLLM:
    '''Wrapper for LLM chain.'''

    def __init__(self, verbose=True, model="gpt-4", additional_tools=None, enable_terminal=False):
        '''Initialize the LLM chain.'''
        self.chain = get_chat_chain(verbose, model, additional_tools=additional_tools,
                                    enable_terminal=enable_terminal)

    def run(self, instructions):
        '''Run the LLM chain.'''
        try:
            result = self.chain.run(instructions)
            return result
        except Exception as e:
            # Workaround for issue https://github.com/hwchase17/langchain/issues/1358.
            if "Could not parse LLM output:" in str(e):
                return str(e).split("Could not parse LLM output:")[1]
            else:
                raise e


def get_chat_chain(verbose=True, model="gpt-4", additional_tools=None,
                   agent=AgentType.CHAT_ZERO_SHOT_REACT_DESCRIPTION,
                   enable_terminal=False, max_iterations=30,
                   max_tokens=None):
    '''Initialize the LLM chain with useful tools.'''
    if os.getenv("OPENAI_API_TYPE") == "azure" or (os.getenv("OPENAI_API_BASE") is not None and "azure" in os.getenv("OPENAI_API_BASE")):
        engine = model.replace(".", "")
        llm = ChatOpenAI(model=model, max_tokens=max_tokens,
                         model_kwargs={"engine": engine})
    else:
        llm = ChatOpenAI(model=model, max_tokens=max_tokens)

    default_tools = ["human", "requests_get"]
    if enable_terminal:
        default_tools += ["terminal"]
    tools = load_tools(default_tools, llm)

    tools += [
        Tool(
            name="python",
            func=PythonREPLTool().run,
            description="Useful for executing Python code with Kubernetes Python SDK client. Results should be print out by calling `print(...)`. Input: Python code. Output: the result from the Python code's print()."
        )
    ]

    if os.getenv("GOOGLE_API_KEY") and os.getenv("GOOGLE_CSE_ID"):
        tools += [
            Tool(
                name="Search",
                func=GoogleSearchAPIWrapper(
                    k=3,
                    google_api_key=os.getenv("GOOGLE_API_KEY"),
                    google_cse_id=os.getenv("GOOGLE_CSE_ID"),
                ).run,
                description="search the web for current events or current state of the world"
            )
        ]

    if additional_tools is not None:
        tools += additional_tools

    memory = ConversationSummaryBufferMemory(
        llm=llm,
        memory_key="chat_history",
        return_messages=True)
    chain = initialize_agent(
        tools, llm, agent=agent, memory=memory,
        agent_kwargs={"output_parser": ChatOutputParser()},
        verbose=verbose, max_iterations=max_iterations,
        handle_parsing_error=handle_parsing_error)
    return chain


def handle_parsing_error(error) -> str:
    '''Helper function to handle parsing errors from LLM.'''
    # Workaround for issue https://github.com/hwchase17/langchain/issues/1358.
    response = str(error).split("Could not parse LLM output:")[1].strip()
    if not response.startswith('```'):
        response = response.removeprefix('`')
    if not response.endswith('```'):
        response = response.removesuffix('`')
    return response
