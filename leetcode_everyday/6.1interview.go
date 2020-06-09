package main

import (
	"fmt"
	"strings"
)

func LongestStr(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	length := len(strs[0])
	tmp := ""
	longest := "" //最长公共前缀
	for n := 0; n < len(strs[0]); n++ {
		tmp = string(strs[0][n])
		fmt.Println(tmp)
		for _, s := range strs {
			// 不含有或者长度超过首个字符串 return
			fmt.Println("s:", s)
			b := strings.HasPrefix(s, longest)
			// fmt.Println(b)
			// fmt.Println("s.len", len(s))
			bigger := len(s) > length
			// fmt.Println("length", length)
			if bigger || !b {
				// fmt.Println("不和条件")
				return longest
			} else {
				// fmt.Println("可以继续")

			}
		}
		longest += tmp
	}
	return longest
}
func main() {
	s := []string{"abeeeeef", "ab", "ab"}
	fmt.Print(LongestStr(s))
}
