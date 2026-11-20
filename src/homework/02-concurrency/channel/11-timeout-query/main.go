// 题目：综合练习·超时查询——query 在 delay 之后把结果写进通道，wait 用 select 让「结果」与
// time.After(timeout) 赛跑。用 100 毫秒超时分别查 40 毫秒的「快查询」和 300 毫秒的「慢查询」：
// 前者拿到结果，后者返回超时——验证超时控制写在调用方这边，不必给查询本身设截止时间。

package main

import (
	"fmt"
	"time"
)

// query 模拟一次异步查询：delay 之后把结果写进通道，并立刻返回该通道
func query(name string, delay time.Duration) <-chan string {
	ch := make(chan string, 1) // 带缓冲：超时后协程仍能写完退出，不会永久阻塞在发送上
	go func() {
		time.Sleep(delay)
		ch <- name + " 的结果"
	}()
	return ch
}

// wait 等结果，超过 timeout 就当超时；ok 为 false 表示没等到
func wait(ch <-chan string, timeout time.Duration) (string, bool) {
	select {
	case v := <-ch:
		return v, true
	case <-time.After(timeout): // 到时后这个通道会收到一个值，select 就走到这一支
		return "", false
	}
}

func main() {
	if v, ok := wait(query("快查询", 40*time.Millisecond), 100*time.Millisecond); ok {
		fmt.Println("拿到结果：", v)
	} else {
		fmt.Println("超时：快查询没有在截止时间内返回")
	}

	if v, ok := wait(query("慢查询", 300*time.Millisecond), 100*time.Millisecond); ok {
		fmt.Println("拿到结果：", v)
	} else {
		fmt.Println("超时：慢查询没有在截止时间内返回")
	}
}
