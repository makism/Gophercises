package main

import (
	"golang.org/x/net/html"
	"strings"
)

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
