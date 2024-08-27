---
layout: assignment
due: 2024-09-03 23:59:59 -0800
github_url: https://classroom.github.com/a/JuYmgyy4
published: true
---

## Requirements

1. You will develop and test an HTML cleaner in Go.
1. Your implementation must compile and pass *your* tests `TestExtract` and `TestCleanHref`
1. Before you add files to your repo, do a `$ go clean` so you don't add/commit build products like executables
1. Your test will run automatically in your repo as a GitHub action.

You will be given mocked HTML data and you should **extract** all the words and urls from the HTML and **clean** 
the urls. Input would look something like this:

```html
<!DOCTYPE html>
<html>
    <head>
        <title>CS272 | Welcome</title>
    </head>
    <body>
        <p>Hello World!</p>
        <p>Welcome to <a href="https://cs272-f24.github.io/">CS272</a>!</p>
    </body>
</html>
```

**extract:** You will extract all of the words between html tags and the urls within the href Tag Attributes.

To tackle this problem, you will need to `go get "golang.org/x/net/html"` and use `html.Parser`  to 
comb through the HTML parse nodes and their content. See the package docs [here](https://pkg.go.dev/golang.org/x/net@v0.12.0/html)
and see the Parser and its functions [here](https://pkg.go.dev/golang.org/x/net/html#Parse). Your 
**extract** function should produce output as follows:


```go
body := []byte(`
<body>
  <p>Some text here</p>
  <a href="http://example.com">Example</a>
</body>
`)

words, hrefs := extract(body)
// words == ["Some", "text", "here", "Example"]
// hrefs == ["http://example.com"]
```

**clean:** You will clean the urls found within the extract function.

Your extract function will produce urls. Some of which will be relative urls. For example, on our course
website, the link to our syllabus is `<a href="/syllabus/">Syllabus<a>`. We want to remember that this is off
of a base url (in this case, our course site: `https://cs272-f24.github.io/`).

```go
host := "https://cs272-f24.github.io/"
hrefs := []string{ "/", "/help/", "/syllabus/", "https://gobyexample.com/" }

for _, url := range hrefs {
  fmt.Println(clean(host, href))
}

/*
  output:
https://cs272-f24.github.io/
https://cs272-f24.github.io/help/
https://cs272-f24.github.io/syllabus/
https://gobyexample.com/
*/
```

## Given

Recall the testing example we covered in class [here](https://github.com/cs272-f24/inclass/tree/main/week02).

## Rubric
Your lab will receive the score indicated by the GitHub autograding action using this rubric. Use these names exactly, including case.
1. 70 points: TestExtract (do you extract the correct words and hrefs).
1. 30 points: TestCleanHref (do you clean the hrefs accurately)
