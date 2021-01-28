package main

import "fmt"

func main() {
	n, m := 0, 0 //n个人m个团队
	fmt.Scan(&n, &m)
	memenber := make([][]int, m)
	total := 0
	//0号是否知道？
	isHeadKnown := false
	for i := 0; i < m; i++ {
		//第i组含有的人数为t
		t := -1
		fmt.Scan(&t)

		memenber[i] = make([]int, t)
		for j := 0; i < t; i++ {
			no := 0
			fmt.Scan(&no)
			memenber[i][j] = no
			if t == 0 {
				isHeadKnown = true
			}
			total++
		}
	}
	// 0号不在队伍里面
	if !isHeadKnown {
		fmt.Println(0)
	}
	fmt.Println(total)

}

func isContained(men []int, i int) bool {
	for _, v := range men {
		if v == i {
			return true
		}
	}
	return false
}
