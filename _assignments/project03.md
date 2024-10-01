---
layout: assignment
due: 2024-10-07 23:59:59 -0800
github_url: https://classroom.github.com/a/IhcE_aVz
published: true
---

## Requirements

1. You will evolve your project03 code to use a persistent [SQLite](https://sqlite.org/index.html) database rather than an in-memory map of maps
1. Your solution must continue to support previous test cases
1. Your solution must not introduce [SQL Injection](https://go.dev/doc/database/sql-injection) vulnerabilities
1. Your solution must use a Go `interface` to build an abstraction for the persistent SQLite database vs. the in-memory map of maps
    1. Make the indexing code switchable on the command line using a Go [flag](https://pkg.go.dev/flag), e.g. 
        ```sh
        $ go run . -index=inmem
        ``` 
        
        or 

        ```sh
        $ go run . -index=sqlite
        ```


## Given

1. We will [discuss](/slides/sql-in-go.html) several approaches to program SQLite databases using Go
1. We will demonstrate programming SQL using the [raw APIs from mattn](https://github.com/mattn/go-sqlite3) and the [ORM APIs from jinzhu](https://gorm.io/)

## Rubric
1. 5 pts: `TestExtract()`
1. 5 pts: `TestCleanHref()`
1. 5 pts: `TestDownload()`
1. 5 pts: `TestCrawl()`
1. 5 pts: `TestSearch()`
1. 15 pts: `TestStop()`
1. 40 pts: `TestTfIdf()`
1. 20 pts: Mandatory code review. 
    1. You will demonstrate completion of the requirements during your code review (no new test cases).
    1. Please sign up for a different grader than you met for project01 or project02.

## Implementation tips
1. You'll need SQL `pragma foreign_keys = on;` for each database connection.
1. If you have SQL `UNIQUE` columns, you don't need to add an index for those because SQLite does so automatically. 
1. However, if you want to `SELECT` for columns in a table, SQLite needs to have an index for each searchable column. In lecture, we described that the `hits` (or `frequency`) table contains non-unique columns that you might want to `SELECT` for
1. If you want to search the SQL index using a join, here's a tip for how to start that query `SELECT urls.name, hits.count FROM urls`. You can finish it.