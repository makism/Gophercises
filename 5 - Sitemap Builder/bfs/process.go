package main

import (
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func parseUrl(url string) []string {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return nil
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil
	}

	links := make([]string, 0, 100)
	processNode(doc, &links)

	return links
}

func processNode(n *html.Node, links *[]string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		newLink := extractLink(n, "https://gophercises.com")

		if newLink != "" {
			*links = append(*links, newLink)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processNode(c, links)
	}
}

func extractLink(n *html.Node, url string) string {
	for _, v := range n.Attr {
		if v.Key == "href" {
			//if strings.HasPrefix(v.Val, url) {
			//if strings.HasPrefix(v.Val, "./") ||
			//	strings.HasPrefix(v.Val, "../") ||
			//	strings.HasPrefix(v.Val, "/") ||
			//	(v.Val == "/") == false ||
			//	(strings.HasPrefix(v.Val, "http://") == false && strings.HasPrefix(v.Val, "https://") == false) {
			if strings.HasPrefix(v.Val, "/") {
				return url + v.Val
			}
		}
	}

	return ""
}
