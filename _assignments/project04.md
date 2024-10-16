---
layout: assignment
due: 2024-10-24 23:59:59 -0800
github_url: https://classroom.github.com/a/U6iSSLew
published: true
---

## Requirements

1. Features
    1. For this project, you will crawl `www.usfca.edu` and produce a searchable index
    1. Your web server will produce clickable search results by generating `<a>` tags using the `<title>` of each web page
1. Ethical crawling
    1. You will use a `Crawl-delay` of 100 milliseconds
    1. You will honor the `Disallow` rules in USF's `robots.txt`
    1. Hopefully we won't get rate-limited!
1. Go concurrency
    1. You will adapt the components of your project03 `sqlite` (not `inmem`) crawler to run concurrently
    1. You will use at least two goroutines and communicate between them using Go channels

## Given

1. We will discuss potential approaches in lecture
1. I strongly recommend watching [Rob Pike's talk](https://www.youtube.com/watch?v=f6kdp27TYZs) on Go's concurrency features

## Rubric

1. 50 pts. TestTfIdf
1. 50 pts. Code review for concurrency and clickable search results
1. For code written by a Large Language Model, you will receive a zero.