package tool

import (
	"fmt"

	set "gopkg.in/fatih/set.v0"
)

//两个数组交集
func CheckWhoNotWork(done []string, all []string) {
	has := set.New(set.ThreadSafe)
	allP := set.New(set.ThreadSafe)
	for _, i := range done {
		has.Add(i)
	}
	for _, i := range all {
		allP.Add(i)
	}

	//并集
	// unionSet := set.Union(a, b)
	// fmt.Printf("并集:%v\n", unionSet)

	//交集
	// intersectionSet := set.Intersection(a, b)
	// fmt.Printf("交集:%v\n", intersectionSet)

	//差集
	// diffS1S2 := set.Difference(a, b)
	// fmt.Printf("差集(属a不属b):%v\n", diffS1S2)

	diffS2S1 := set.Difference(allP, has)
	fmt.Printf("差集(属b不属a):%v\n", diffS2S1)

}
