// 题目：写 setZero(n int)，在函数内把 n 设为 0。在 main 里传入 42，调用前后各打印一次，确认外面仍是 42。

package main

import "fmt"

func setZero(n int) {
	n = 0
}

func SetZero() {
	num := 42
	fmt.Println(num)

	setZero(num)

	fmt.Println(num) // 仍是 42：setZero 改的是副本
}
