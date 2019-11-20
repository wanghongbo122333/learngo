package goRoutine

import (
	"fmt"
	"time"
)

func goRoutineRun() {
	go func() {
		for {
			say("hello")
		}
	}()
	say("world")
	time.Sleep(1)
	fmt.Println("over")
	// go spinner(100 * time.Millisecond)
	// const n = 45
	// fibN := fib(n) // slow
	// fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func say(s string) {
	fmt.Println(s)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
