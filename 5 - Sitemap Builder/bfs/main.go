package main

import (
	"flag"
	"fmt"
	"os"
)

type Link struct {
	Text  string `json:"Text"`
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
		Text:  "",
		Depth: 0,
		To:    make([]Link, 0, 5),
	}
}

var opts = NewOptions()

func main() {
	flag.StringVar(&opts.Url, "url", "https://gophercises.com", "URL to crawl.")
	flag.IntVar(&opts.Depth, "depth", 5, "Maximum link depth.")
	flag.Parse()

	if opts.Url == "" {
		flag.Usage()
		os.Exit(1)
	}

	sitemap := bfs(opts.Url, opts.Depth)

	fmt.Println("final sitemap: ", sitemap)
}
