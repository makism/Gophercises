package main

import (
	"golang.org/x/net/html"
	"strings"
)

func processNode(n *html.Node, l *Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		newLink := extractLink(n, l.Href)

		if newLink != "" {
			nl := NewLink()
			nl.Depth = l.Depth + 1
			nl.Href = l.Href + newLink
			nl.Text = newLink

			l.To = append(l.To, nl)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processNode(c, l)
	}
}

func extractLink(n *html.Node, url string) string {
	for _, v := range n.Attr {
		if v.Key == "href" {
			if strings.HasPrefix(v.Val, "./") ||
				strings.HasPrefix(v.Val, "../") ||
				strings.HasPrefix(v.Val, "/") ||
				(strings.HasPrefix(v.Val, "http://") == false && strings.HasPrefix(v.Val, "https://") == false) {
				return v.Val
			}
		}
	}

	return ""
}
