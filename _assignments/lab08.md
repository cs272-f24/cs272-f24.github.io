---
layout: assignment
due: 2024-12-03 23:59:59 -0800
github_url: https://classroom.github.com/a/WOARv_96
published: true
---

## Background

1. For this assignment, you will use [aider](https://aider.chat) to build a web UI on top of your project06 solution
1. We will build this project during the lecture meeting on Tuesday December 3rd
1. Since aider is a python program which uses [Anthropic Claude Sonnet](https://www.anthropic.com/claude/sonnet), you'll need to set those up using the instructions below.

## Environment setup

1. Accept the lab08 assignment and clone the repo
1. Copy your project06 code into your lab08 repo
1. Check the version of python you have on your laptop using `python --version` or `python3 --version`. If it's between 3.09 and 3.12, you're good, skip to the next step
    1. If you have python 3.13, aider [unfortunately](https://github.com/Aider-AI/aider/issues/1984) does not work with python 3.13, so you need to set up a virtual python environment. There are many ways to do this, but one easy way is to use [conda](https://docs.conda.io/projects/conda/en/stable/user-guide/getting-started.html)
    1. You should download Miniconda onto your laptop and run the installer
    1. After conda has been installed you can create a new python 3.12 environment using `conda create -n env312 python=3.12` where `env312` is any name you choose for the name of the virtual python environment
    1. Then run `conda activate env312` and you should see `(env312)` on the left side of your shell prompt
1. Install aider using the instructions on their web page
1. Create an [Anthropic](https://www.anthropic.com/) API key and put it in your environment, just like you did for OpenAI. You'll need to give them a payment method, but it should cost less than $1 to develop this lab.

## Objective

1. You will use aider to add a web server (HTTP `Handle`, `ListenAndServe` etc.) to your project06 solution
  - Aider can replace your input loop (`io.Scanner` etc.) if you wish. You don't need to keep both
  - Aider can develop a web page which can accept user prompts and chat with the LLM
  - Aider can reuse your chat code to call OpenAI to answer the same kinds of questions that your project06 solution answers.
1. Feel free to use your own creativity to prompt Aider to make an attractive web app. Some ideas:
  - Appearance similar to other chat apps
  - Light/dark mode
  - Spinner while waiting for OpenAI
  - Theme in the style of ... 
1. You should do as little manual coding and fixing as possible. Prompt aider to do the work.

## Rubric

1. 80 pts - Something runs
1. 20 pts - Attractive and useful
