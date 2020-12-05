package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	dbFilename:= *flag.String("db", "mapping.db", "Mapping database.")
	dbInit := false //*flag.Bool("init", true, "Initialize the dataase; load in the mappings from the file \"input.yml\" into a bucket.")
	flag.Parse()

	if _, err := os.Stat(dbFilename); err != nil {
		dbInit = true
	}

	db := DbOpenOrCreate(dbFilename)
	if dbInit {
		parsedYaml, err := parseYAML("input.yml")

		if err != nil {
			fmt.Println("Failed parsing the given yaml file.")
			os.Exit(1)
		}

		DbInit(db, parsedYaml)
	}

	defer db.Close()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	dbHandler, err := DbHandler(db, mapHandler)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func parseYAML(filename string) ([]yamlRecords, error) {
	data, err := ioutil.ReadFile(filename)

	var resultYAML []yamlRecords
	err = yaml.Unmarshal([]byte(data), &resultYAML)
	return resultYAML, err
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
