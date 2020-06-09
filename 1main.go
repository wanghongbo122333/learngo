package main

// import (
// 	"fmt"

// 	"learngo/command"
// 	"learngo/flag"
// 	"learngo/goroutine"
// )

// func modifySlice1(data []int) {
// 	data = nil
// }
// func modifySlice2(data []int) {
// 	data[0] = 0
// }

// func main() {
// 	a := []int{1, 2}
// 	fmt.Println(a)
// 	modifySlice1(a)
// 	fmt.Println(a)
// 	modifySlice2(a)
// 	fmt.Println(a)
// }
// func main() {
// 	runtime.GOMAXPROCS(1)
// 	int_chan := make(chan int, 1)
// 	string_chan := make(chan string, 1)
// 	int_chan <- 1
// 	string_chan <- "hello"
// 	select {
// 	case value := <-int_chan:
// 		fmt.Println(value)
// 	case value := <-string_chan:
// 		panic(value)
// 	}
// }
func hammingDistance(x, y int) int {
	n := x ^ y
	res := 0
	for n != 0 {
		if n & 1 {
			res++
		}
		n /= 2
	}

	return res
}
