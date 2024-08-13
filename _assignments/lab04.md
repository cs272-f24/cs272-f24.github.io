---
layout: assignment
due: 2023-09-25 23:59:59 -0800
github_url: https://classroom.github.com/a/aHaeieFc
published: false
---

## Requirements
1. In this lab you will build your own test case library by splitting some books 
(represented in HTML) into chapters that you will later search.
1. The class web site contains copies of [10 of the most-downloaded books](https://cs272-0304-f23.github.io/tests/top10/) from
Project Gutenberg. The HTML follows a predictable pattern that you will
parse using Go's `html.Parser`. You do NOT have to crawl that page. You can simply use the URLs.

## Given
1. The key piece of information is that every chapter in the given books is demarcated
by the HTML element `<div class="chapter"...>`. 
1. When you parse the file, you will find a `<head>` element which should go in 
every chapter you create
1. Every chapter you create should end with the HTML elements `</body></html>`
1. We will discuss strategies for building data structures and algorithms to process the input files
1. We will discuss how you might use `html.Render()`, `html.Node.RemoveChild()`, `filepath`, `os.Create()`, and `os.MkdirAll()`

## Rubric
100 pts. for `TestSplit()` as shown by the autograder web page. 

## Testing
1. Your `TestSplit()` should use [this hash function](/tests/lab04/hashFile.txt) to generate a `uint32` hash digest 
1. The expected hash values for the files are given [here](/tests/lab04/hashes.txt)
1. To provide a reference example, [this file](https://github.com/cs272-0304-f23/lab04-given/blob/main/chap01.html) contains the correct contents for Chapter 1 of Frankenstein, and produces the hash value `668812877`. Use View Source in your browser to see how the the generated file contains:
    1. `<html>`
    1. The contents of the `<head>` node
    1. `<body>`
    1. The contents of the `<div class=chapter>` node
    1. `</body></html>`
