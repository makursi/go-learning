// 题目：无缓冲通道上先在 main 里发送、再起接收协程——运行即 deadlock（all goroutines are asleep）。
// 故意做成本来就会死锁的顺序：直接 go run，看到的 deadlock 就是本题要观察的现象。

package main

import "fmt"

func main() {
	ch := make(chan int) // 无缓冲通道

	ch <- 10 // 此刻还没有接收方就绪，这里永久阻塞

	go func() {
		val := <-ch

		fmt.Println("已接收数据", val)
	}()

	fmt.Println("永远打印不到这里")
}
