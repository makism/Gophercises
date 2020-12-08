package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const DefaultStoryJson = "gopher.json"

func main() {
	story, _ := LoadStory(DefaultStoryJson)

	start := story["intro"]
	presentChapter(story, start)
}

func presentChapter(story Story, chapter Chapter) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	l := len(chapter.Title)
	maxOptions := len(chapter.Options)

	fmt.Println(chapter.Title)
	fmt.Println(strings.Repeat("=", l))

	for _, s := range chapter.Story {
		fmt.Println(s)
	}

	fmt.Println()

	if maxOptions > 0 {
		fmt.Println("Responses:")
		for i, s := range chapter.Options {
			fmt.Printf("[%d] %s\n", i+1, s.Text)
		}

		for {
			char, _, err := keyboard.GetSingleKey()
			if (err != nil) {
				panic(err)
			}

			defer func() {
				_ = keyboard.Close()
			}()

			tmp := string(char)

			if tmp == "q" {
				fmt.Println("Exiting...")
				os.Exit(0)
			}

			index, _ := strconv.Atoi(tmp)

			if index == 0 || index >= maxOptions+1 {
				continue
			}

			nextChapter := story[chapter.Options[index-1].Arc]
			presentChapter(story, nextChapter)
		}
	} else {
		fmt.Println("Fin.")
		os.Exit(0)
	}
}
