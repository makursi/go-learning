// 题目：无缓冲交接：写函数 double(n int, ch chan int)，把 n * 2 发送到通道；在 main 里 go double(21, ch)，再接收并打印（应是 42）。

package main

import "fmt"

func double(n int, ch chan int) {
	ch <- n * 2
}

func UnbufferedDocking() {
	ch := make(chan int)

	go double(21, ch)

	val := <-ch

	fmt.Println(val)
}
