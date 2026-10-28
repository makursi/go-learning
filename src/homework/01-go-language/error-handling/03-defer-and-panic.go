// 题目：注册三个 defer 打印 A、B、C，再打印一句「函数体」，看输出顺序是不是「函数体 → C → B → A」；
// 然后用 defer + recover 接住 mustPositive 在非正数时的 panic，并验证 panic 之后的那句「其余逻辑」不会执行。

package main

import "fmt"

func mustPositive(n int) {
	if n <= 0 {
		panic("错误")
	}
}

func DeferAndPanic() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")

	fmt.Println("函数体")

	// 这个 defer 后注册，所以先于 A、B、C 执行——panic 被它接住后，
	// 本函数剩下的语句不再执行，直接按后进先出跑完其余 defer。
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("捕获到程序错误:", info)
		}
	}()

	mustPositive(-1)

	fmt.Println("其余逻辑")
}
