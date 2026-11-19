// 题目：两个协程分别延迟 100 微秒与 200 微秒后往各自的通道写入，用 select 取先到达的那个并打印——体会 select 是「谁先就绪就走谁」。

package main

import (
	"fmt"
	"time"
)

func workerA(ch chan string) {
	time.Sleep(100 * time.Microsecond)
	ch <- "A 完成"
}

func workerB(ch chan string) {
	time.Sleep(200 * time.Microsecond)
	ch <- "B 完成"
}

func SelectFastest() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go workerA(ch1)
	go workerB(ch2)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}
