---
layout: assignment
due: 2023-09-16 23:59:59 -0800
github_url: https://classroom.github.com/a/J-dTww1O
published: true
---
## Requirements
1. You will build a web crawler which incorporates the elements of the prior lab assignments, including their test cases: 
    1. Download from a URL
    1. Extract non-markup text using Go's `html.Parser`
    1. Extract `href` attributes from `<a>` tags
1. Lab03 adds the following new features, which must include new test cases
    1. An in-memory inverted index based on Go's `map`
    1. An HTTP server which can search an index
    1. [Snowball stemmer](https://github.com/kljensen/snowball)
1. Your solution will crawl the text of [Romeo and Juliet](/tests/rnj/), which I got from [Project Gutenberg](https://www.gutenberg.org/) and divided
into an HTML document per scene.
1. You will write a `TestSearch` function which tests these three cases for the inverted index, returning the term frequency for each one (i.e. the map of string to frequency)
    1. Search `https://cs272-f24.github.io/tests/rnj/sceneI_30.0.html` for `Verona` should find one hit:
        ```
        "https://cs272-f24.github.io/tests/rnj/sceneI_30.0.html": 1
        ```
    1. Search `https://cs272-f24.github.io/tests/rnj/sceneI_30.1.html` for `Benvolio` should return 26 hits:
        ```
        "https://cs272-f24.github.io/tests/rnj/sceneI_30.1.html": 26
        ```
    1. Search `https://cs272-f24.github.io/tests/rnj/` for `Romeo` should return these hits:
        ```
        "https://cs272-f24.github.io/tests/rnj/sceneI_30.0.html":  2,
        "https://cs272-f24.github.io/tests/rnj/sceneI_30.1.html":  22,
        "https://cs272-f24.github.io/tests/rnj/sceneI_30.3.html":  2,
        "https://cs272-f24.github.io/tests/rnj/sceneI_30.4.html":  17,
        "https://cs272-f24.github.io/tests/rnj/sceneI_30.5.html":  15,
        "https://cs272-f24.github.io/tests/rnj/sceneII_30.2.html": 42,
        "https://cs272-f24.github.io/tests/rnj/":                  200,
        "https://cs272-f24.github.io/tests/rnj/sceneI_30.2.html":  15,
        "https://cs272-f24.github.io/tests/rnj/sceneII_30.0.html": 3,
        "https://cs272-f24.github.io/tests/rnj/sceneII_30.1.html": 10,
        "https://cs272-f24.github.io/tests/rnj/sceneII_30.3.html": 13,
        ```
1. You can use `reflect.DeepEqual()` to test the equivalence of your `got` map and your `want` map

## Given

In lecture meetings we will discuss and provide sample code for:
1. Hash tables, the underlying technology for go `map`
1. Inverted indexes, which can be used to facilitate searching a corpus of documents
1. Word stemming

## Rubric
1. 100 pts: correctness as shown by the autograder workflow on github.com
    1. 5 pts each for `TestExtract`, `TestCleanHref`, `TestDownload`, and `TestCrawl` continuing to pass
    1. 80 pts for `TestSearch` passing
1. NB: grading happens on the branch named `dev`