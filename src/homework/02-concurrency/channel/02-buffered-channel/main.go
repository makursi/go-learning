// 题目：定义容量为 2 的缓冲通道，子协程依次发 10、20，主协程再逐个接收并打印——体会「发送方不必等接收方」与无缓冲的差别。

package main

import "fmt"

func sendTwo(ch chan int) {
	ch <- 10
	ch <- 20
}

func main() {
	ch := make(chan int, 2) // 有缓冲：容量 2，前两次发送不会阻塞

	go sendTwo(ch)

	val1 := <-ch
	val2 := <-ch

	fmt.Println(val1)
	fmt.Println(val2)
}
