package leetcode

// 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-the-duplicate-number
func findDuplicate(nums []int) int {
	var m = make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			return v
		}
	}
	return -1
}

// 二分法
func findDuplicate2(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		// fmt.Print("mid", mid)
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}

// func main() {
// 	var array []int
// 	array = append(array, 1, 2, 3, 4, 5, 1)
// 	fmt.Print(findDuplicate2([]int{1, 2, 3, 4, 5, 1}))
// }
