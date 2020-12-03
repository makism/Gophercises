package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type parsedArgs struct {
	csv   string
	limit int
}

type question struct {
	problem string
	answer  int
}

func prepareQuestions(pa parsedArgs) []question {
	var questions []question
	csvFile, err := os.Open(pa.csv)

	if err == nil {
		r := csv.NewReader(csvFile)

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			convRecord, err := strconv.Atoi(record[1])

			if err == nil {
				q := question{
					record[0],
					convRecord,
				}

				questions = append(questions, q)
			}
		}
	} else {
		fmt.Println("Failed reading the CSV file.")
		os.Exit(1)
	}

	return questions
}

func main() {
	pa := parsedArgs{}
	flag.StringVar(&pa.csv, "input", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.IntVar(&pa.limit, "limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	questions := prepareQuestions(pa)

	timerDone := time.NewTimer(time.Duration(pa.limit) * time.Second)

	correctAnswers := 0

mainLoop:
	for i, q := range questions {
		chanAnswer := make(chan string)
		go func() {

			fmt.Printf(
				"Problem #%d: %s = ",
				i+1,
				q.problem,
			)

			var answer string
			fmt.Scanf("%s", &answer)

			chanAnswer <- answer
		}()

		select {
		case <-timerDone.C:
			fmt.Println("\nTime's up!")
			break mainLoop
		case readAnswer := <-chanAnswer:
			checkAnswer, _ := strconv.Atoi(readAnswer)

			if checkAnswer == q.answer {
				correctAnswers++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(questions))

}
