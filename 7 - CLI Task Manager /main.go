package main

import (
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		showUsage()
	}

	dbInit := false
	dbFilename:= "todo.db"

	if _, err := os.Stat(dbFilename); err != nil {
		dbInit = true
	}

	Db = DbOpenOrCreate(dbFilename)
	if dbInit {
		DbInit()
	}

	intercept(args[1:])
}
