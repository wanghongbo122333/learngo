package twritter

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func twritterRun() {
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	//args:格式化目标	最小单元长度 tab字符宽度 计算tab宽度时会加上 用于填充的字符 格式化控制字符
	// w.Init(os.Stdout, 0, 8, 0, '\t', 2)
	// str := "\tab\tc\td\t."
	// str2 := "123\t12345\t1234567\t123456789\t."
	// // fmt.Println("原始输出是：", str)
	// // amount, err := fmt.Fprintln(w, str)
	// // fmt.Printf("返回值是： %d %v\n", amount, err)
	// fmt.Fprintln(w, str)
	// w.Flush()
	// fmt.Fprint(w, str2)
	// w.Flush()

	// Format right-aligned in space-separated columns of minimal width 5
	// and at least one blank of padding (so wider column entries do not
	// touch each other).
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()
}
