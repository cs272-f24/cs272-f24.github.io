---
layout: assignment
due: 2023-11-14 23:59:59 -0800
github_url: https://classroom.github.com/a/aAAmCJtJ
published: false
---

## Requirements

1. You will evolve your project04 solution to build SQLite indexes using `sitemap.xml`
1. You may choose any web site which has a `sitemap` in its `robots.txt`. Here are some examples:
    1. usfca.edu
    1. usd.edu
    1. scu.edu
    1. smc.edu
    1. gonzaga.edu (recursive)
    1. ucsc.edu (recursive)
    1. dvc.edu
    1. apple.com (multiple)
    1. google.com (recursive)
    1. nvidia.com (by language)
    1. intel.com (recursive)
    1. openai.com
    1. pkg.go.dev
    1. cnn.com
    1. npr.org

1. You will develop two test cases using TF-IDF scores similar to those provided for project03 and lab05.
1. Limits
    1. If no `crawl-delay` is found, you may use 100ms as a default
    1. Since some sites have more documents than others, you may stop crawling after 10,000 documents

## Given

1. We will discuss the XML format and Go's support for XML in lecture
1. You'll want to use the Go tags to assist with unmarshalling the XML
1. Here's an example Go data structure which can be used to unmarshal the XML given at [http://usfca.edu/sitemap.xml](https://www.usfca.edu/sitemap.xml) (use View Source to the the raw XML)
    ```go
    type(
    	Sitemap struct {
            Loc     string `xml:"loc"`
            LastMod string `xml:"lastmod"`
        }

        SitemapIndex struct {
            Sitemaps []Sitemap `xml:"sitemap"`
        }
    )
    ```
1. You can extend this to the `<urlset>` and `<url>` structures as shown at [https://www.usfca.edu/sitemap.xml?page=1](https://www.usfca.edu/sitemap.xml?page=1) (use View Source to see the raw XML)

## Rubric

1. 50 pts for `TestCase1` tested locally (not on Github)
1. 50 pts for `TestCase2` tested locally (not on Github)
1. Provide a screen shot in your Github repo which shows the test cases passing. Crawling a real web site will take too long for the 5 minute time limit for Github Actions
1. 5 pts Extra Credit. Choose a site which has a recursive `sitemap.xml`