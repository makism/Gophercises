package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

type Link struct {
	Href  string `json:"Href"`
	Depth int    `json:"Depth"`
	To    []Link `json:"To"`
}

type Options struct {
	Url   string
	depth int
}

func NewOptions() Options {
	return Options{
		Url:   "",
		depth: 10,
	}
}

func NewLink() Link {
	return Link{
		Href:  "",
		Depth: 0,
		To:    make([]Link, 0, 5),
	}
}

func main() {
	opts := NewOptions()

	flag.StringVar(&opts.Url, "url", "https://golang.org", "URL to crawl.")
	flag.IntVar(&opts.depth, "depth", 10, "Maximum link depth.")
	flag.Parse()

	if opts.Url == "" {
		flag.Usage()
		os.Exit(1)
	}

	if resp, err := http.Get(opts.Url); err == nil {
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		root := NewLink()
		processNode(doc, &root)
	}
}

func prettyPrint(d Link) {
	if res, err := json.MarshalIndent(d, "", "    "); err == nil {
		fmt.Println(string(res))
	}
}

func processNode(n *html.Node, l *Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		newLink := extractLink(n)

		if newLink != "" {
			if l.Href == "" {
				l.Href = newLink
			} else {
				nl := NewLink()
				nl.Depth = l.Depth + 1
				nl.Href = newLink
				l.To = append(l.To, nl)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processNode(c, l)
	}
}

func extractLink(n *html.Node) string {
	for _, v := range n.Attr {
		if v.Key == "href" {
			if strings.HasPrefix(v.Val, "/") {
				return v.Val
			}
		}
	}

	return ""
}
