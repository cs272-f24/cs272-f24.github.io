---
layout: assignment
due: 2023-10-30 23:59:59 -0800
github_url: https://classroom.github.com/a/1qzVcG6f
published: false
---

## Background

1. For this lab, you will design a database schema using [SQLite](https://sqlite.org/index.html)
1. In a subsequent project, we will adapt your project03 code to use this schema 


## Tools setup

1. For this lab you will need to install the `sqlite3` command-line interpreter
1. Check if you already have `sqlite3`  using `$ which sqlite3`. 
1. If you don't have it you can install it in one of two ways:
    1. MacOS users can use homebrew `% brew install sqlite`
    1. From the [SQLite web site](https://sqlite.org/download.html)

## Requirements

1. Using the lecture discussion about relational modeling, develop a schema for your project03 data structures and represent it in two ways:
1. Sketch an ER diagram on paper using your entities and cardinality
1. Create tables using the `sqlite3` command-line interpreter and output the results using
    ```text
    sqlite> .output lab06.sql
    sqlite> .dump
    ```
1. Please submit a photo of your sketch and lab06.sql to your assignment repo

## Rubric

1. 100 pts for credible effort