package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	//"golang.org/x/net/html"
)

type options struct {
	htmlFile string
	url string
	localMode bool
}

type Link map[string] string

func main() {
	opts := options{localMode: true}

	flag.StringVar(&opts.htmlFile, "html", "", "A local HTML file to parse.")
	flag.StringVar(&opts.url, "url", "", "A URL to parse.")
	flag.Parse()

	if opts.htmlFile == "" && opts.url == "" {
		flag.Usage()
		os.Exit(1)
	}

	if opts.htmlFile != "" {
		if fp, err := os.Open(opts.htmlFile); err == nil {
			process(fp)
			defer fp.Close()
		} else {
			fmt.Println(err)
		}
		os.Exit(0)
	}

	if opts.url != "" {
		if resp, err := http.Get(opts.url); err == nil {
			process(resp.Body)
			defer resp.Body.Close()
		} else {
			fmt.Println(err)
		}
		os.Exit(0)
	}

}
