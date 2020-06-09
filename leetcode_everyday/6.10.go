package main

import "strconv"

//判断是否是回文数
func isPalindrome(x int) bool {
	//负数或者以0结尾的数都false
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	src := strconv.Itoa(x)
	for i := 0; i < len(src); i++ {
		if src[i] == src[len(src)-i-1] {
			continue
		} else {
			return false
		}
	}
	return true
}
