package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

func process(data io.Reader) {
	doc, err := html.Parse(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	processNode(doc)
}

func processNode(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		if l := extractLink(n); l != nil {
			fmt.Println(l)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processNode(c)
	}
}

func extractLink(n *html.Node) Link {
	l := Link{}

	if n.FirstChild.Type == html.TextNode {
		l["Text"] = strings.TrimSpace(extractText(n))
	}

	for _, v := range n.Attr {
		if v.Key == "href" {
			l["Href"] = v.Val

			return l
		}
	}

	return nil
}

func extractText(n *html.Node) string {
	text := ""

	for c:= n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text = text + c.Data
		} else {
			text = text + extractText(c)
		}
	}

	return text
}
