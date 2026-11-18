// 题目：启动 3 个协程，内部分别 Sleep 50ms / 200ms / 100ms 后打印；主协程只 Sleep 80ms 就打印「结束」。
// 观察哪些输出丢了，并说明为什么多协程不该靠自己定延时。

package main

import (
	"fmt"
	"time"
)

// 答：无法把各协程的执行时间与主程序的执行时间精确匹配，主程序结束时还没跑完的
// 协程会被直接丢弃，于是输出时有时无。要等所有协程，应该用 WaitGroup 而不是 Sleep。
func makeGoroutine(sleepTime time.Duration, name string) {
	time.Sleep(sleepTime)

	fmt.Println("协程", name, "开始执行")
}

func WhyNotRelyOnSleep() {
	go makeGoroutine(50*time.Millisecond, "A")
	go makeGoroutine(200*time.Millisecond, "B")
	go makeGoroutine(100*time.Millisecond, "C")

	time.Sleep(80 * time.Millisecond)

	fmt.Println("结束") // B 与 C 还没有结果：主协程先退出了
}
