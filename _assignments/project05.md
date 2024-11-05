---
layout: assignment
due: 2024-11-12 23:59:59 -0800
github_url: https://classroom.github.com/a/PuZC1TB3
published: true
---

## Requirements

For this project you will evolve your vector database solution from lab06 by adding:

1. A feature to ask questions of an LLM using command-line input
1. A solution for the "PHILosophy" problem
1. A prompting strategy which answers questions about courses by department, course name, instructor, times, and location

## Given

In lecture, we will discuss:

1. How to use Go to read input from the command line
1. The problems of relying on vector similarity to build prompts
1. Opportunities to improve the database implementation using metadata
1. Opportunities to improve the prompts and leverage the capabilities of the LLM

## Rubric

1. 50 pts. - correctness including required test cases
    1. `TestPhil` to prompt the LLM to correctly answer the question, "What courses is Phil Peterson teaching in Fall 2024?"
    1. `TestPHIL` to prompt the LLM to correctly answer the question, "Which philosophy courses are offered this semester?"
    1. `TestBio` to prompt the LLM to correctly answer the question, "Where does Bioinformatics meet?"
    1. `TestGuitar` to prompt the LLM to correctly answer the question, "Can I learn guitar this semester?"
    1. `TestMultiple` to prompt the LLM to correctly answer the question, "I would like to take a Rhetoric course from Phil Choong. What can I take?"
1. 50 pts. - code review
    1. You must build an abstraction around the chat completion API so you can reuse it
    1. You must build the prompt in a general way, without hard-coding the desired LLM output in the prompt