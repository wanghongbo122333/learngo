package main

import (
	"fmt"
)

func main() {

	var n int
	var ans int

	fmt.Scan(&n) //n个人
	// ans = resolve(n)
	ans = compute(n)
	//求answer
	fmt.Printf("%d\n", ans)
}
func f(n int64, m int64) int64 {
	var MAX int64 = 1000*1000*1000 + 7
	//a个中取b个
	if m > n {
		return 0
	}
	if m == 0 {
		return 1
	}
	return f(n-1, m-1)%MAX + f(n-1, m)%MAX

}
func resolve(n int64) int64 {

	var res int64 = 0
	var i int64 = 0
	for i = 1; i <= n; i++ {
		res = res + (f(n, i) * i)
	}
	return res
}

func compute(a int) int {
	sum := 0
	k := 0
	for i := a; i > 0; i-- {
		k = 1
		tmp := 1
		for j := a; k <= i; j-- {
			tmp = tmp * j
			k++
		}
		sum += tmp
	}
	return sum
}
