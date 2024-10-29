---
layout: assignment
due: 2024-11-04 23:59:59 -0800
github_url: https://classroom.github.com/a/JbM3fIOC
published: true
---

## Background

1. For this lab and subsequent projects, you will import a corpus of USF course information into a vector database
1. For this lab, we will simply populate the database
1. For future assignments, we will use the database to get chatbot answers which are more accurate than the publicly available LLMs can provide

## Tools

### ChromaDB

ChromaDB is one of the popular vector databases, offering semantic similarity. The easiest way to install ChromaDB on your laptop is using Docker

1. [Docker Desktop](https://www.docker.com/)
1. The template repo contains a `docker-compose.yaml` file which uses Docker Desktop to run the latest version of ChromaDB, putting the database  in `chromadb/`
1. You can run it using `docker compose up`
1. You can program ChromaDB in Go using the [amikos-tech](https://go-client.chromadb.dev/) package

### A Large Language Model (LLM)

You can choose between [OpenAI](https://openai.com/) and [Ollama](https://ollama.com/)

1. To use OpenAI's cloud API, you'll need an OpenAI Project Key
    1. Go to openai.com > Products > API, and sign in. I used Google and my USF credentials
    1. You'll need to set up a payment method. I originally put $20 into it, and have $19.69 remaining, so it's not expensive. The `GPT4oMini` model is good and inexpensive
    1. You'll need access to the project API key in your code, but you should not commit the key to github. A reasonable approach is a shell script and Go's `os.Getenv()`
    1. You can program OpenAI in Go using the [sashabaranov](https://github.com/sashabaranov/go-openai) package

1. If you prefer, you can download one of the Meta/Facebook models to your laptop
    1. Ollama provides a local web server on top of the Meta (aka Facebook) models such as Llama 3.2 without any cost or API key
    1. If you choose this approach make sure you have [enough compute resources](https://llamaimodel.com/requirements-3-2/) 

## Given

The template repo contains

1. A CSV file containing Fall 2024 course information, which was kindly provided by the registrar's office
1. A `.gitignore` file which ignores `chromadb/` and `openai_project_key.sh` so you don't commit them by accident
1. The `docker-compose.yaml` file described above


## Requirements

1. You will develop Go code which reads the CSV file and builds a ChromaDB collection using the fields of the CSV file. 
1. We will discuss several approaches to that in lecture
1. Your code will put all 2,740 CSV rows into the database
1. You will add a test case to insert a known row into the collection and query to get it back
1. (Optional) You could have some fun with the name of the course and instructor if you like

## Rubric

1. TestVectorQuery - 100 pts. Will be graded manually as it takes too long for Github Actions