package main

import (
	"fmt"
	"learngo/leetcode"
)

func main() {
	// s := "a"
	// j := "srasrassaaa"
	// fmt.Println(leetcode.NumJewelsInStones(s, j))
	// m := [][]int{{1, 1, 0}, {1, 1, 0}, {1, 1, 1}, {1, 1, 1}, {0, 0, 0}, {1, 1, 1}, {1, 0, 0}}

	// fmt.Println(leetcode.IsToeplitzMatrix(m))
	// t := {}int{1}
	// fmt.Println(leetcode.Average(t))
	// f := strings.Replace("hello", "e", "r", -1)
	// fmt.Println(f)
	// g := []int{14, 15, 16, 22, 53, 60}
	// // fmt.Println(leetcode.KWeakestRows(m, 6))
	// fmt.Println(leetcode.TwoSum(g, 76))
	// fmt.Println(leetcode.InvalidateIP("192.156.1.5"))
	var l1 leetcode.ListNode
	var l2 leetcode.ListNode
	var l3 leetcode.ListNode

	l1.Val = 1
	l2.Val = 2
	l3.Val = 3

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = nil

	s := leetcode.ReversePrint(&l1)
	fmt.Println(s)

}
