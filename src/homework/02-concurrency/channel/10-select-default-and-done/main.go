// 题目：select 的使用——一次把三种写法走完：
// 1) 带 default：对空的无缓冲通道 select，没有 case 就绪，立刻走 default 而不阻塞；
// 2) 不带 default：等 50 毫秒后才到的 "ping"，select 会阻塞到它真的到达；
// 3) 停止通道：在 for + select 里循环收 3 条消息，读到 done 关闭就退出（数据通道不关闭也能结束）。

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// 1) 有 default：没有任何 case 就绪时立刻执行 default
	select {
	case v := <-ch:
		fmt.Println("读到：", v)
	default:
		fmt.Println("通道为空，立刻走 default")
	}

	// 2) 没有 default：select 一直等，直到某个 case 就绪
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- "ping"
	}()
	select {
	case v := <-ch:
		fmt.Println("阻塞等到：", v)
	}

	// 3) 停止通道：读到 done 就 return，ch 始终没有被 close
	done := make(chan struct{})
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- fmt.Sprintf("消息%d", i)
		}
		close(done) // 结束信号走 done，不靠关闭数据通道
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println("循环收到：", v)
		case <-done:
			fmt.Println("收到 done，退出循环")
			return
		}
	}
}
