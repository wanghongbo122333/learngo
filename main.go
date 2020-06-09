package main

import (
	"fmt"
)

func main() {
	n := 0
	m := 0
	var tree []int
	//一共n个数，被m整除
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		tree = append(tree, x)
	}
	var sum = make([][]int, n, n)
	fmt.Print(sum)
	fmt.Print(n, m)
	fmt.Print(tree)
	// fmt.Println("result:", totalCount(tree, n, m))

	//设置sum[i][i]
	for i := 0; i < n; i++ {
		sum[i][i] = tree[i]
	}
	//设置sum[0][i]
	for i := 0; i < n; i++ {
		sum[0][i] = sumTree(tree, 0, i)
	}
	for i := 1; i < n; i++ {
		for j := i + 1; i < n; j++ {
			sum[i][j] = sum[i][j-1] + tree[j]
		}
	}
	fmt.Println(sum)
}
func totalCount(tree []int, length int, toMod int) int {
	result := 0
	// treeLength := 1 //区间长度最少为1
	//i 区间起点
	for i := 0; i < length; i++ {
		//j为区间长度
		for j := 1; j <= length-i; j++ {
			// fmt.Println("i,j:", i, j)
			tmp := sumTree(tree, i, i+j)
			if tmp%toMod == 0 {
				result++
			}
		}

	}
	return result
}
func sumTree(tree []int, start int, end int) int {
	sum := 0
	for _, i := range tree[start:end] {
		fmt.Print("i:", i)
		sum = sum + i
	}
	fmt.Println("sum = ", sum)
	return sum
}
