// 题目：用无缓冲通道让子协程把 100 交给主协程：起一个协程发送，主协程接收并打印。

package main

import "fmt"

func sendOne(ch chan int) {
	ch <- 100 // 向通道发送数据
}

func main() {
	ch := make(chan int) // 无缓冲：发送与接收必须同时就绪

	go sendOne(ch) // 先起发送方，否则主协程阻塞在 <-ch 时没人送值

	data := <-ch // 主协程在这里等子协程把值送达

	fmt.Println("接收到数据:", data)
}
