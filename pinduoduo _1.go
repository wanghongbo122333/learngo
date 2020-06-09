package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	var array []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		array = append(array, x)
	}
	fmt.Println(array)
	fmt.Println("result:", cal(array, n))
}

func cal(array []int, length int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(array)))
	fmt.Println(array)
	if length <= 2 {
		return sum(array, length)
	}

	var pay []int
	//从大到小，每买两个，消除一个
	for i := 0; i < length; i++ {
		if (i+1)%3 != 0 {
			pay = append(pay, array[i])
		}
	}
	return sum(pay, len(pay))
}
func sum(array []int, length int) int {
	sum := 0
	for i := 0; i < length; i++ {
		sum = sum + array[i]
	}
	return sum
}
