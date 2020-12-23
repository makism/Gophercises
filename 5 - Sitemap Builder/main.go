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
	Depth int
}

func NewOptions() Options {
	return Options{
		Url:   "",
		Depth: 5,
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

func main() {
	flag.StringVar(&opts.Url, "url", "https://golang.org", "URL to crawl.")
	flag.IntVar(&opts.Depth, "depth", 5, "Maximum link depth.")
	flag.Parse()

	if opts.Url == "" {
		flag.Usage()
		os.Exit(1)
	}

	root := parseStartPage(opts.Url)

	wg.Add(1)
	for i := range root.To {
		go parseSubPage(&root, i, true)
	}
	wg.Wait()

	prettyPrint(root)
}

func parseSubPage(root *Link, index int, release bool) {
	url := opts.Url + root.To[index].Href

	if resp, err := http.Get(url); err == nil {
		page, _ := html.Parse(resp.Body)

		newLink := root.To[index]

		if newLink.Depth >= opts.Depth {
			if release {
				wg.Done()
			}
			return
		}

		processNode(page, &newLink)
		root.To[index] = newLink

		for i, _ := range newLink.To {
			parseSubPage(&newLink, i, false)
		}
	}

	if release {
		wg.Done()
	}
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
		root.Href = url
		processNode(doc, &root)

		return root
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	return Link{}
}
