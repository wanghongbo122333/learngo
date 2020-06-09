package main

func subarraySum(nums []int, k int) int {
	record := map[int]int{0: 1}
	l := len(nums)

	p := make([]int, l)
	// 前i个数的和
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			p[i] = nums[0]
		} else {
			p[i] = p[i-1] + nums[i]
		}

	}
	//sum[i,j]=p[j]-p[i]
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			sum[i][j] = p[j] - p[i]
		}

	}
}

//计算从start到end的sum
func sum(array []int, start, end int) int {
	sum := 0
	for i := start; i < end; i++ {
		sum += array[i]
	}
	return sum
}

// 前i个数的和
func sum_i(end int, array []int) []int {
	for i := 0; i < count; i++ {

	}
}
