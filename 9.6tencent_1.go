package main

import "fmt"

func main() {
	n1, n2 := 0, 0
	fmt.Scan(&n1)
	arr1 := make([]int, 0)
	for i := 0; i < n1; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		arr1 = append(arr1, tmp)
	}
	fmt.Scan(&n2)
	arr2 := make([]int, 0)
	for i := 0; i < n2; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		arr2 = append(arr2, tmp)
	}
	for i, j := 0, 0; i < n1 && j < n2; {
		if arr1[i] > arr2[j] {
			i++
			continue
		}
		if arr1[i] < arr2[j] {
			j++
			continue
		}
		if arr1[i] == arr2[j] {
			fmt.Printf("%d ", arr1[i])
			i++
			j++
			continue
		}
	}

}

// idxOfMax := make([]int, 0) //记录最大值更新的位置
// 	//求最长公共子序列
// 	//dp[i][j]表示a1[0..i]中和a2[0..j]中最长公共子串的长度
// 	dp := make([][]int, n1+1)
// 	for i := 0; i < n1; i++ {
// 		dp[i] = make([]int, n2)
// 	}
// 	for i := 0; i < n1; i++ {
// 		for j := 0; j < n2; j++ {
// 			if arr1[i] == arr2[j] {
// 				dp[i+1][j+1] = dp[i][j] + 1
// 			}
// 		}
// 	}
