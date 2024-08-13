---
layout: assignment
due: 2023-12-05 23:59:59 -0800
github_url: https://classroom.github.com/a/1abNhs2a
published: false
---

## Requirements

1. Visual Appeal using CSS
    1. Using raw CSS or one of the popular frameworks such as [Bulma](https://bulma.io/) or [Pure](https://purecss.io/) make your search results more visually appealing. ([Bootstrap](https://getbootstrap.com/) is also ok but more involved)
    1. Use Go's [HTML templates](https://pkg.go.dev/html/template) to substitute your search results into an HTML snippet including style information.
1. Implement one of these Information Retrieval features including two test cases
    1. Auto-complete on page `<title>`s using a [Search Trie](https://en.wikipedia.org/wiki/Trie). Your search results must update as each keystroke is typed.
    1. Search for images using the contents of the `alt` attribute and display the image in search results
    1. Include a "snippet" of text surrounding a search term so users of your search engine can see the context around the term. You can see examples in public search engines.

## Given

In lecture we will discuss:
1. the syntax for HTML templates and their application in Go
1. the structure of CSS selectors, properties, and values
1. how to reference style sheets in your HTTP ResponseWriter

An example of the `alt` attribute: `<img src="dog.jpeg" alt="my dog Goldy playing with her toy">`

An [example of snippets](https://duckduckgo.com/?t=h_&q=computer+science&ia=web) searching duckduckgo.com for "computer science"

## Rubric

Since this project wraps the semester, code review meetings will not include any TODO recommendations

1. Visual appeal
    1. Minimal effort (30 pts)
    1. Good effort (40 pts)
    1. Great effort (50 pts)
1. Implement one of the three IR features given above (50 pts)
1. Extra credit: implement two of the IR features
    1. Good effort (5 pts)
    1. Great effort (5 additional points)
