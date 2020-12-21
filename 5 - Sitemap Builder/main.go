package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"sync"
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

var wg sync.WaitGroup

var opts = NewOptions()

var depth int = 0

func main() {
	flag.StringVar(&opts.Url, "url", "https://golang.org", "URL to crawl.")
	flag.IntVar(&opts.depth, "depth", 10, "Maximum link depth.")
	flag.Parse()

	if opts.Url == "" {
		flag.Usage()
		os.Exit(1)
	}

	root := parseStartPage(opts.Url)

	for i, _ := range root.To {
		wg.Add(1)
		go func(root *Link, index int) {
			url := opts.Url + root.To[index].Href

			if resp, err := http.Get(url); err == nil {
				page, _ := html.Parse(resp.Body)

				newLink := NewLink()
				newLink.Depth = root.To[index].Depth + 1

				processNode(page, &newLink)

				root.To[index].To = append(root.To[index].To, newLink)
			}

			wg.Done()
		}(&root, i)
	}

	wg.Wait()

	prettyPrint(root)
}

func parseStartPage(url string) Link {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err == nil {
		doc, err := html.Parse(resp.Body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		root := NewLink()
		processNode(doc, &root)

		return root
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	return Link{}
}
