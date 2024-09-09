---
layout: assignment
due: 2024-09-09 23:59:59 -0800
github_url: https://classroom.github.com/a/X5DJINEw
published: true
---

## Requirements
Project01 adds support for http

1. You will develop and test a web crawler, reusing your `clean()` and `extract()` code from lab02
1. Your implementation must compile and pass *your* tests `TestDownload` and `TestCrawl`
1. Your implementation will use the published test data
  - [`https://cs272-f24.github.io/tests/project01/index.html`](/tests/project01/index.html)
  - [`https://cs272-f24.github.io/tests/project01/href.html`](/tests/project01/href.html)
  - [`https://cs272-f24.github.io/tests/project01/simple.html`](/tests/project01/simple.html)
  - [`https://cs272-f24.github.io/tests/project01/style.html`](/tests/project01/style.html)
`
1. Your test will run automatically in your repo as a GitHub action.

Given a base URL, download the web page from that address using Go's [`net/http`](https://pkg.go.dev/net/http) package. We can then leverage the 
extraction functionality from `lab02` to get the words and cleaned URLs in this downloaded page. The next 
step is the stem the words we get from this webpage.

**download** Given a url, download the content using the `http` package.
```go
url := "https://cs272-f24.github.io/"
body, err := download(url)

// now you can extract(body) if there are no errors!
```

**crawl** Given a seed URL, **download** the webpage, **extract** the words and URLs, and add all **cleaned URLs** 
found to a download queue. This queue should only crawl each url **once** and repeat this process for each url.
```go
url := "https://cs272-f24.github.io/"
crawl(url)

/*
  your output might look something like this
  although this is not part of the automated test:

  download: url=https://cs272-f24.github.io/
  download: result=ok
  download: url=https://cs272-f24.github.io/help/
  download: result=ok
  download: url=https://cs272-f24.github.io/syllabus/
  download: result=ok
  download: url=https://cs272-f24.github.io/syllabus/assignments/lab01.html
  download: result=ok
*/
```
## Given

In lecture, we will demonstrate:
1. the [`net/http`](https://pkg.go.dev/net/http) client
1. [`net/http/httptest`](https://pkg.go.dev/net/http/httptest) for "mocking" a server
1. Simple web page [test cases](https://cs272-f24.github.io/tests/project01/). You can copy the HTML source into your test code.

## Rubric
Your lab will receive the score indicated by the GitHub autograding action using this rubric:
1. 5 pts: `TestExtract()`
1. 5 pts: `TestCleanHref()`
1. 35 pts: `TestDownload()` (do you download the correct content)
1. 35 pts: `TestCrawl()` (can you crawl a given seed url and download it along with all of the links embedded within)
1. 20 pts: 1:1 code review to review your coding style and test cases
1. NB: grading happens on the branch named `dev`