---
layout: assignment
due: 2023-11-06 23:59:59 -0800
github_url: https://classroom.github.com/a/SO7hdwM0
published: false
---

## Requirements

1. You will evolve your project03 code to use a persistent [SQLite](https://sqlite.org/index.html) database rather than an in-memory map of maps
1. Your solution must continue to support recent feature additions, such as concurrency and `robots.txt`
1. Your solution must protect your database from concurrent writes
1. Your solution must not introduce [SQL Injection](https://go.dev/doc/database/sql-injection) vulnerabilities

## Given

1. We will [discuss](/slides/sql-in-go.html) several approaches to program SQLite databases using Go
1. We will use the same TF-IDF [test cases](/tests/lab05/test-cases.go) from lab05 (that is, disallowing `chap31.html`) and [corpus](/tests/project03/top10.zip) from project03
1. In order to complete crawling in less than 30 seconds as required by `go test`, a new [`robots.txt `](/tests/project04/robots.txt) is given, which omits `crawl-delay`

## Rubric

1. 10 pts. TestDisallow (same results as the in-memory map of maps)
1. 50 pts. TestTfIdf
1. 40 pts. Code review for SQL and relational design. Please pass the test cases before coming to the code review meeting.
1. 5 pts. (Extra credit)
    1. Use a Go `interface` to build an abstraction for the persistent SQLite database vs. the in-memory map of maps
    1. Make the indexing code switchable on the command line using a Go [flag](https://pkg.go.dev/flag), e.g. 
        ```sh
        $ go run . -index=inmem
        ``` 
        
        or 

        ```sh
        $ go run . -index=sqlite
        ```

## Implementation tips
1. You'll need SQL `pragma foreign_keys = on;` for each database connection.
1. If you have SQL `UNIQUE` columns, you don't need to add an index for those because SQLite does so automatically. 
1. However, if you want to `SELECT` for columns in a table, SQLite needs to have an index for each searchable column. In lecture, we described that the `hits` (or `frequency`) table contains non-unique columns that you might want to `SELECT` for
1. If you want to search the SQL index using a join, here's a tip for how to start that query `SELECT urls.name, hits.count FROM urls`. You can finish it.