// 题目：写函数 hello(name string) 打印「你好, xxx」；用 go 调用它，并用 time.Sleep 短暂等待后结束（体会「协程能跑起来」即可）。

package main

import (
	"fmt"
	"time"
)

func hello(name string) {
	fmt.Println("你好", name)
}

func StartGoroutine() {
	go hello("choria")

	// 先 go 再等：Sleep 放在 go 之前的话，主协程会在新协程有机会运行前就结束。
	// 这里只是凑合演示，正式代码用 WaitGroup 等。
	time.Sleep(1 * time.Second)

	fmt.Println("主协程执行完毕")
}
