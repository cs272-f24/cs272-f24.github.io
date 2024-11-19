---
layout: assignment
due: 2024-11-25 23:59:59 -0800
github_url: https://classroom.github.com/a/p07JjOEG
published: true
---

## Requirements

In this project, you will enhance your tool-calling lab to run in a loop, accumulating a slice of
`openai.ChatCompletionMessage` which provides a running session (or "context") for your questions

1. Add context

```text
Search> Who is teaching CS 272?
Philip Peterson is teaching CS 272
Search> What's his email address?
Philip Peterson's email address is phpeterson@usfca.edu
```

1. Write an original tool. Think of a useful scenario for tool calling and implement it

```
Search> Take me to the web page where I can ...
```

1. Handle multiple argument and multiple tools

```text
Search> What are Phil Peterson and Greg Benson teaching?

Search> What courses is Phil Peterson teaching in LS G12?
```

## Given

1. We will discuss design strategies for utilizing the LLM's context window
1. You will probably need to reuse some of your project05 code which provides hints to the LLM in a user prompt

## Rubric

100 pts. - Code review on Tuesday. Will be graded during the meeting. No Github pull requests or reviewers.