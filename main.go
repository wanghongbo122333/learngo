package main

import (
	"fmt"

	"learngo/command"
	"learngo/flag"
	"learngo/goroutine"
)

func main() {
	fmt.Println()
	// a := ""
	var b string
	fmt.Println("" == b)
	command.CommandRun()
	flag.FlagRun()
	goroutine.GoRoutineRun()
}
