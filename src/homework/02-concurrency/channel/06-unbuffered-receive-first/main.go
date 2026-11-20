// 题目：无缓冲通道上先起接收协程、再由 main 发送——程序正常结束，说明无缓冲为什么要求两边配合。（这是死锁版的修改版，见 07-unbuffered-send-first.go）

package main

import "fmt"

func main() {
	ch := make(chan int) // 无缓冲通道

	// 先起接收方，让它就绪并阻塞在 <-ch 上
	go func() {
		val := <-ch

		fmt.Println("已接收数据", val)
	}()

	ch <- 10 // 接收方已就绪，发送不会阻塞

	fmt.Println("发送完成")
}
