// 题目：无缓冲通道上先在 main 里发送、再起接收协程——运行即 deadlock（all goroutines are asleep）。
// 故意不接进 main 调度；要观察死锁，把 main 里的调度临时改成 UnbufferedSendFirst() 再 go run。

package main

import "fmt"

func UnbufferedSendFirst() {
	ch := make(chan int) // 无缓冲通道

	ch <- 10 // 此刻还没有接收方就绪，这里永久阻塞

	go func() {
		val := <-ch

		fmt.Println("已接收数据", val)
	}()

	fmt.Println("永远打印不到这里")
}
