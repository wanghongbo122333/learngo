package main

import "fmt"

// 给定数组nums,求数组res，res[i]为nums所有乘积除nums[i]外
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = multiple(nums, i)
	}
	return res
}
func multiple(nums []int, j int) int {
	res := 1
	for i := 0; i < len(nums); i++ {

		if i != j {
			if nums[i] == 0 {
				return 0
			}
			res *= nums[i]
		}
	}
	return res
}

//前后缀相乘
func productExceptSelf2(nums []int) []int {
	length := len(nums)
	//L[i]表示num[i]的前缀乘积 R[i]表示nums[i]的后缀乘积
	L, R, answer := make([]int, length), make([]int, length), make([]int, length)
	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[1+i]
	}
	for i := 0; i < length; i++ {
		answer[i] = L[i] * R[i]
	}
	return answer
}

//前后缀相乘 缩小额外空间
func productExceptSelf3(nums []int) []int {
	length := len(nums)
	//answer[i]存放nums[i]的前缀积
	answer := make([]int, length)
	answer[0] = 1

	for i := 1; i < length; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	//直接乘以前缀得到结果
	//然后得到后缀积R:=0
	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i]
	}

	return answer
}
func main() {
	fmt.Println(productExceptSelf3([]int{1, 2, 3, 4}))
}
