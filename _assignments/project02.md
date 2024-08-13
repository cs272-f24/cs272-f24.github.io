---
layout: assignment
due: 2023-10-03 23:59:59 -0800
github_url: https://classroom.github.com/a/WaZWvs1a
published: false
---
## Requirements
1. For this project, you will extend the functionality of the [MVP](https://en.wikipedia.org/wiki/Minimum_viable_product) search engine you build in project01
1. You will apply these [Information Retrieval (IR)](https://en.wikipedia.org/wiki/Information_retrieval) techniques as discussed in lecture.
    1. Words will again be stemmed using the Snowball stemmer
    1. Your inverted index will omit Stop Words using this list
    1. You will rank search results using the TF-IDF technique
1. You will reuse these elements of prior assignments:
    1. the corpus of 10 books you built for lab04
    1. the crawler and inverted index you built for project01
1. Architectural requirements:
    1. Indexing and searching must be separated in your code so that your web server runs while crawling is working
    1. Your local web server (on `localhost` or `127.0.0.1`) will serve the top10 corpus for your local crawler
    1. You must not use global variables. Make at least one element of your code use methods in an object-oriented style (hint: an index might be helpful)

## Given
1. From previous assignments:
    1. [Snowball stemmer](https://github.com/kljensen/snowball)
    1. Corpus of [10 books from Project Gutenberg](https://cs272-0304-f23.github.io/tests/top10/)
1. Lecture material on scoring and ranking search results
1. You can use the Go [`sort.Interface`](https://pkg.go.dev/sort) interface to do the sorting
1. You can use the [stopwords-iso](https://github.com/stopwords-iso/stopwords-en) collection of English stopwords. The `json` list is formatted similarly to Go.

## Rubric
1. 5 pts: `TestDownload()`
1. 5 pts: `TestCrawl()`
1. 10 pts: `TestSearch()`
1. 20 pts: `TestStopWords()`
1. 30 pts: `TestTfIdf()`
1. 20 pts: Code review. Please, no enormous files or functions. Break your code and test cases into smaller files for readability.

## Required Test Cases
1. Data structures for test cases are provided:
    - [Stopword tests](/tests/project02/stop_tests.go)
    - [TF-IDF tests](/tests/project02/tfidf_tests.go)
1. You'll see that I called the lab04 corpus "top10". You can rename that if you like, but don't change the wanted URLs or scores.

## Implementation Tips
1. You'll want to use the Go `float64` type to calculate the TF-IDF score. For example, you might calculate  
    ```go
    tf := (float64) termCount / (float64) wordsInDoc
    ```
1. Some terms generate the same TF-IDF score. In order to make the test cases deterministic (same order for all hits with the same score), your implementation of the `Less()` function in `sort.Interface` should check for the same score, and if so, also sort by the URL. For example:
    ```go
    if scoreOfI == scoreOfJ {
        return urlOfJ > urlOfI
    }
    ```
1. Remember to push your code to the `dev` branch for grading
1. Remember to push your corpus so the autograder can test against the chapter files you created
1. You do NOT need to include the code for lab04 or implement the `TestSplit` test case