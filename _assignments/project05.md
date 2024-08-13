---
layout: assignment
due: 2023-11-21 23:59:59 -0800
github_url: https://classroom.github.com/a/5aZFQ0pV
published: false
---

## Requirements

1. You will evolve your lab07 solution, again crawling the web site you chose in lab07, and adding several new Information Retrieval features:
1. Wildcard searching
    1. Your index.html page will include new UI to do wildcard searching. You have two options:
        - Add a checkbox to the form for wildcards search
        - Add a character to the search term (such as '*') to indicate wildcard search
    1. Your SQLite query will add support for `LIKE` and `%`
1. Word N-grams
    1. Your crawler will add support for bigram (2-gram) word sequences to support searches like "soccer team" or "computer science"
1. Page titles
    1. You will use the HTML parser to extract the content of `title` tags and present those as search hits rather than the URLs
    1. You will add the page title to your SQL relational model and your search handler logic

## Given

1. We will discuss indexing strategies and end-to-end design approaches

## Rubric

1. 50 pts. - correctness as shown by `go test` on your local machine. You will construct test cases for 
    1. a 2-gram search expressing hits using page titles (25 pts) 
    1. a wildcard search expressing hits using page titles (25 pts)
1. 50 pts. - code review

## Implementation tips
1. Here is an example HTML form which includes a checkbox
    ```html
    <html>
        <body>
            <form action="/search" method="get">
                <label for="inputBox">Search</label>
                <input id="inputBox" name="term"/>
                <button type="submit">Go</button>
                <br>
                <input type="checkbox" id="wildcard" name="wildcard" value="wildcard">
                <label for="wildcard">Wildcard Search</label>
            </form>
        </body>
    </html>
```
1. You can use SQLite [`LIKE`](https://www.sqlitetutorial.net/sqlite-like/) to do wildcard searches
1. You can add a new SQL table for bigram hits