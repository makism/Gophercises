package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const DEFAULT_CSV string = "problems.csv"
const DEFAULT_LIMIT int = 30

type settings struct {
	csv string
	limit int
}

func showUsage(binary string) {
	fmt.Print("Usage ./"+ binary + "\n" +
		"\t-input string\n" +
		"\t       a csv file in the format of 'question,answer' (default \"problems.csv\")\n" +
		"\t-limit int\n" +
		"\t       the time limit for the quiz in seconds (default 30)\n")
}

func parseArgs(args []string) settings {
	sets := settings{
		limit: DEFAULT_LIMIT,
		csv: DEFAULT_CSV,
	}

	for i, v := range args {
		if v == "-input" {
			sets.csv = args[i + 1]
		} else if v == "-limit" {
			sets.limit, _ = strconv.Atoi(args[i + 1])
		}
	}

	return sets
}

func main() {
	args := os.Args[1:]

	if l := len(args); l == 0 {
		showUsage(os.Args[0])
		os.Exit(1)
	}

	parsed_args := parseArgs(args)
	fmt.Println(parsed_args)

	csv_file, _ := os.Open(parsed_args.csv)
	r := csv.NewReader(csv_file)

	correct_answers := 0

	i := 0
	for {
		i += 1

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"Problem #%d: %s = ",
			i,
			record[0],
		)

		var answer string
		fmt.Scanf("%s", &answer)

		conv_answer, _ := strconv.Atoi(answer)
		conv_record, _ := strconv.Atoi(record[1])
		if conv_answer == conv_record {
			correct_answers +=1
		}
	}

	fmt.Println("You scored ", correct_answers, " out of ", i, ".")

}
