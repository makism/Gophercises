package main

import "fmt"

func main() {
	userInput := make(chan string)

	go func(){
		fmt.Printf("Please give camelcase string: ")

		var answer string
		fmt.Scanf("%s", &answer)

		userInput <- answer
	}()


	select {
	case answer := <-userInput:

		ans := []byte(answer)

		if startsWithLowercase(ans) {
			firstWord, rest := extractFirstWord(ans)

			fmt.Printf("%s\n", firstWord)

			for {
				word, rest2, ok := extractWord(rest)

				if ok == false {
					break
				}

				rest = rest2

				fmt.Printf("%s\n", word)
			}
		} else {
			fmt.Println("This doesn't look like a camel-case string...")
		}
	}
}

func extractFirstWord(b []byte) ([]byte, []byte) {
	endIndex :=0

	counter := 0
	for i, v := range b {
		if isUppercase(v) {
			endIndex = i
			break
		} else {
			counter++
		}
	}

	if counter == len(b) {
		return b, nil
	}

	firstWord := b[0:endIndex]
	rest := b[endIndex:]

	return firstWord, rest
}

func extractWord(b []byte) ([]byte, []byte, bool) {
	endIndex := 0
	ok := true

	if len(b) == 0 {
		return nil, nil, false
	}

	if startsWithUppercase(b) {
		counter := 1
		for i, v := range b[1:] {
			if isLowercase(v) {
				counter++
				continue
			} else {
				endIndex = i
				break
			}
		}

		if counter == len(b) {
			return b, nil,true
		}
	} else {
		ok = false
	}

	return b[0:endIndex + 1], b[endIndex + 1:], ok
}

