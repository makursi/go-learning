// 题目：把上一题改成 sync.WaitGroup：三个任务都 Done 后再打印「结束」，确认每次都能看到全部输出。

package main

import (
	"fmt"
	"sync"
	"time"
)

func makeWaitGroupExample(sleepTime time.Duration, name string, wg *sync.WaitGroup) {
	defer wg.Done() // 任务完成后，通知 WaitGroup 计数器减一

	time.Sleep(sleepTime)

	fmt.Println("协程", name, "开始执行")
}

func WaitGroupBasics() {
	var wg sync.WaitGroup

	wg.Add(3)

	go makeWaitGroupExample(50*time.Millisecond, "A", &wg)
	go makeWaitGroupExample(200*time.Millisecond, "B", &wg)
	go makeWaitGroupExample(150*time.Millisecond, "C", &wg)

	wg.Wait()

	fmt.Println("所有任务结束")
}
