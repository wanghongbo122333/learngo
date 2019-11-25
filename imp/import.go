package imp

import (
	"fmt"
)

var Global = "hello"

/*
当main函数引入imp这个包的时候，会自动执行init函数，执行顺序是从上到下
init函数有一下几点特性:
init函数在main执行之前，自动被调用执行的，不能显示调用
每个包的init函数在包被引用时，自动被调用
每个包可以有多个init函数
同一个文件中可定义多个init()函数

*/
func init() {
	fmt.Println("自动执行function1")
}
func init() {
	fmt.Println("自动执行function2")
}
