// 题目：写一个 max(a, b int) int，返回较大的那个数；调用并打印结果。

package main

import "fmt"

// 题目要的就是叫 max：它会遮蔽 Go 1.21 起的内建 max，但在本包里够用。
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Max() {
	maxNum := max(10, 20)
	fmt.Println(maxNum)
}
