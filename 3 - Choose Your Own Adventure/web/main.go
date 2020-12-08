package main

import (
	"log"
	"net/http"
)

const DefaultStoryJson = "gopher.json"

func main() {
	story, _ := LoadStory(DefaultStoryJson)

	h := StoryHandler(story)
	http.ListenAndServe(":8080", h)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

