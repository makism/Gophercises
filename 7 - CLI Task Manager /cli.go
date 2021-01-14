package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var validCommands map[string]CommandInterface

func intercept(commandWithParams []string) {
	validCommands = make(map[string]CommandInterface)
	validCommands["add"] = NewCommandAdd()
	//validCommands["do"] = NewCommandDo()
	validCommands["list"] = NewCommandList()

	command := commandWithParams[0]
	if cmd, ok := validCommands[command]; ok {
		cmd.Execute(commandWithParams[1:])
	}
}

func showUsage() {
	str :=`{BIN} is a CLI for managing your TODOs.

Usage:
  {BIN} [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "{BIN} [command] --help" for more information about a command.`

	bin, _ := os.Executable()
	bin = filepath.Base(bin)

	r := strings.NewReplacer("{BIN}", bin)
	fmt.Println(r.Replace(str))
	os.Exit(0)
}
