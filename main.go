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
	// var b string
	// fmt.Println("" == b)//true

	command.CommandRun()
	flag.FlagRun()
	goroutine.GoRoutineRun()
	leetcode.NumJewelsInStones("", "")
}
