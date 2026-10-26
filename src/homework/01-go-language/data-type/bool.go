package main

import "fmt"

func Bool() {
	//布尔：声明一个 bool 变量表示「是否在线」，初始值为 true，打印后将其改为 false 再打印一次。

	online := true
	fmt.Println(online)
	online = false
	fmt.Println(online)
}
