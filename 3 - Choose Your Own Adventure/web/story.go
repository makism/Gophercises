package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Story map[string] Chapter

type Chapter struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc string `json:"arc"`
	} `json:"options"`
}

func LoadStory(jsonFilename string) (Story, error) {
	var stories Story

	data, err := ioutil.ReadFile(jsonFilename)

	if err == nil {
		err = json.Unmarshal(data, &stories)
		if err == nil {
			return stories, nil
		}
		log.Println(err)
	} else {
		log.Println(err)
	}

	return nil, err
}
