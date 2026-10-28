// 题目：写一个 swap(a, b int) (int, int)，交换两个整数并返回；调用后打印结果。

package main

import "fmt"

func swap(a, b int) (int, int) {
	var temp int
	temp = a
	a = b
	b = temp

	return a, b
}

func Swap() {
	fmt.Println(swap(10, 20))
}
