---
layout: assignment
due: 2023-10-27 23:59:59 -0800
github_url: https://classroom.github.com/a/Qjtk9WuF
published: false
---

## Requirements

1. Your web crawler will read a well-formatted `robots.txt` file at the root of your `top10` corpus
1. The `robots.txt` file will contain `User-Agent` records, `Disallow` records, and `Crawl-delay` records
1. If not provided, the default `Crawl-delay` is 500 milliseconds per hostname. That is, your crawler may issue an `http.Get()` to a hostname at most twice per second.

## Given

1. I crawled `http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/index.html` and searched for "blood"
1. The top hit is chap21.html
1. Example `robots.txt` including `disallow: *chap21.html` is given [here](/tests/lab05/robots.txt)
1. Search hits (minus chap21.html) and a sketch for `TestCrawlDelay()` are given [here](/tests/lab05/test-cases.go).

## Rubric

1. 50 pts. `TestCrawlDelay`
1. 50 pts. `TestDisallow`

## Implementation Notes

1. You should evolve your project03 Concurrent Crawler solution, adding a new data structure for the records in a `robots.txt` file. 
1. You should implement the policies for `disallow` rules and `crawl-delay` rules before calling `http.Get()` in your crawler
1. You should use `regexp.MatchString()` to read the records in the `robots.txt` file
1. You should use `Time.Sub()` to check that the resulting `time.Duration` is >= the `crawl-delay` rule
