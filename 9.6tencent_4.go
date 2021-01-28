package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)
	arr := make([]int, n+1)
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		arr[i]=tmp
		b[i]=tmp
	}
	sort.Ints(arr)
	mid:=n/2
	for i := 1; i <=n; i++ {
		if b[i]<=arr[mid] {
			fmt.Println(arr[mid+1])
		}else{
			fmt.Println(arr[mid])
		}
	}
}
