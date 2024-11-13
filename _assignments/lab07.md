---
layout: assignment
due: 2024-11-19 23:59:59 -0800
github_url: https://classroom.github.com/a/I1pr6opA
published: true
---

## Requirements

For this lab, you will evolve your project05 solution as follows:

1. For Named Entity Recognition, you will adopt "structured output" using a JSON schema, if you didn't already do that for project05
1. You will remove the RAG-style hints, which were developed as input to ChatCompletion hints
1. Rather, you will passing the user's question directly to the LLM and provide hints in the style of "tool/function calling"
1. You will add "tools" to the ChatCompletion request
1. Your tool implementation will query ChromaDB and provide structured query results to the LLM
1. The LLM will use the results of the tool call to answer the user's question

## Given

1. We will discuss the fundamentals of JSON schema and tool calling in lecture

## Draft Rubric - subject to change

1. 100 pts - Improve test cases from project05?