---
layout: assignment
due: 2024-09-23 23:59:59 -0800
github_url: https://classroom.github.com/a/gmDMp-5S
published: true
---
## Requirements
1. For this project, you will create an [MVP](https://en.wikipedia.org/wiki/Minimum_viable_product) search engine you build in project01
1. You will apply these [Information Retrieval (IR)](https://en.wikipedia.org/wiki/Information_retrieval) techniques as discussed in lecture.
    1. Words will again be stemmed using the Snowball stemmer
    1. Your inverted index will omit Stop Words using this list
    1. You will rank search results using the TF-IDF technique
1. You will reuse the crawler and inverted index you built for previous assignments
1. Architectural requirements:
    1. Indexing and searching must be separated in your code so that your web server runs while crawling is working
    1. Your local web server (on `localhost` or `127.0.0.1`) will serve the top10 corpus for your local crawler
    1. You must not use global variables. 
    1. Break your code and test cases into smaller files for readability.
    1. Use Go methods with receivers to express your code in an object-oriented way.

## Given
1. From previous assignments:
    1. [Snowball stemmer](https://github.com/kljensen/snowball)
    1. Corpus of [10 books from Project Gutenberg](/top10/)
1. Lecture material on scoring and ranking search results
1. You can use the Go [`sort.Interface`](https://pkg.go.dev/sort) interface to do the sorting
1. You can use the [stopwords-iso](https://github.com/stopwords-iso/stopwords-en) collection of English stopwords. The `json` list is formatted similarly to Go.

## Rubric
1. 5 pts: `TestExtract()`
1. 5 pts: `TestCleanHref()`
1. 5 pts: `TestDownload()`
1. 5 pts: `TestCrawl()`
1. 5 pts: `TestSearch()`
1. 15 pts: `TestStop()`
1. 40 pts: `TestTfIdf()`
1. 20 pts: Code review. Mandatory, as usual. Please sign up for a different grader than you met for project01.

## Required Test Cases
1. You will develop a handful of good test cases for stop words and TF-IDF rankings. 
1. Your test cases will use the "slice of structs" approach demonstrated in lecture to express the expected output
1. You'll see that I called the test corpus "top10". You can rename that if you like

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
1. Remember to push your corpus so the autograder can test against the chapter files 
