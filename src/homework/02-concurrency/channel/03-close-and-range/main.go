// 题目：往容量 3 的通道里发 1、2、3，然后 close，用 range 读出全部值——体会 close 之后 range 会自然结束，而不是卡住。

package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) // 没有这一句，range 会一直等下去（并在最后 panic: all goroutines are asleep）

	for v := range ch {
		fmt.Println(v)
	}
}
