package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	if len(candies) <= 1 {
		return []bool{true}
	}
	max := candies[0]
	for _, i := range candies {
		if i > max {
			max = i
		}
	}
	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			res[i] = true
		}
	}
	return res
}

func main() {
	m := []int{4, 2, 1, 1, 2}
	ex := 1
	fmt.Println(kidsWithCandies(m, ex))

}
