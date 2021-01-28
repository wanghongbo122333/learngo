package main

import (
	"fmt"
	"sort"
)

type node struct {
	val   string
	count int
}

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	res := topKStr(arr, k)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%s %d", res[i].val, res[i].count)
		// fmt.Println(res[i].val, " ", res[i].count)
	}
}
func topKStr(str []string, k int) []node {
	var res []node
	val2Count := make(map[string]int)
	for _, v := range str {
		val2Count[v]++
	}
	var nodeList []node
	for k, v := range val2Count {
		nodeList = append(nodeList, node{k, v})
	}
	//前k多排序
	sort.Slice(nodeList, func(i, j int) bool {
		if nodeList[i].count > nodeList[j].count {
			return true
		} else if nodeList[i].count == nodeList[j].count && nodeList[i].val < nodeList[j].val {
			return true
		}
		return false
	})
	total := len(val2Count)
	max := min(k, total) //如果总数少于k
	for i := 0; i < max; i++ {
		//前k多的
		res = append(res, nodeList[i])
	}
	//前k少的排序
	sort.Slice(nodeList, func(i, j int) bool {
		if nodeList[i].count < nodeList[j].count {
			return true
		} else if nodeList[i].count == nodeList[j].count && nodeList[i].val < nodeList[j].val {
			return true
		}
		return false
	})

	for i := 0; i < max; i++ {
		//前k少的
		res = append(res, nodeList[i])
	}

	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
