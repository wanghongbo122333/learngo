package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n-1 == 0 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = Max(dp[0], dp[1])
	for i := 2; i < n; i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	array := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(rob(array))
}
