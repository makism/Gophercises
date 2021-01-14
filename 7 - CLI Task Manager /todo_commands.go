package main

import (
	"fmt"
	"strings"
)

type CommandInterface interface {
	Execute([]string)
	Help()
}

type CommandMeta struct {
	name string
	help string
	requiresArgs bool 
}

type CommandAdd struct {
	meta CommandMeta
	arg []string
}

type CommandList struct {
	meta CommandMeta
	arg []string
}

type CommandDo struct {
	meta CommandMeta
	arg []string
}

func invokeHelp(args []string, c CommandInterface) {
	if len(args) == 1 && args[0] == "--help" {
		c.Help()
	}
}

func NewCommandList() CommandList {
	return CommandList{meta: CommandMeta{"list", "List all of your incomplete tasks", false}}
}

func (c CommandList) Execute(args []string) {
	invokeHelp(args, c)
	DbListAll()
}

func (c CommandList) Help() {
	fmt.Println(c.meta.help)
}

func NewCommandAdd() CommandAdd {
	return CommandAdd{meta: CommandMeta{"add", "Add a new task to your TODO list", true}}
}

func (c CommandAdd) Execute(args []string) {
	invokeHelp(args, c)

	// Join the string array into one.
	item := strings.Join(args[0:]," ")

	DbAddRecord(item)
}

func (c CommandAdd) Help() {
	fmt.Println(c.meta.help)
}
