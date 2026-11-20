// 题目：创建容量为 2 的 chan string，连发 "A"、"B" 后再发第三个而先不接收——会发生什么？
// 答：会发生死锁，因为通道容量已满，第三个发送永久阻塞。
// 它是一道独立程序：直接 go run，看到的 deadlock 就是本题要观察的现象。

package main

import "fmt"

func main() {
	ch := make(chan string, 2) // 容量 2 的缓冲通道

	ch <- "A"
	ch <- "B"

	ch <- "C" // 缓冲已满，这里永久阻塞 → all goroutines are asleep

	str1 := <-ch
	str2 := <-ch

	fmt.Println(str1)
	fmt.Println(str2)
}
