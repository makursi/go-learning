// 题目：用匿名函数实现「打印 1 到 5」，立刻执行；再把同一个逻辑赋给变量后调用一次。

package main

import "fmt"

func AnonymousFunc() {
	// 写法一：立刻执行（IIFE）
	func() {
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
		}
	}()

	// 写法二：先赋给变量，再调用
	printOneToFive := func() {
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
		}
	}
	printOneToFive()
}
