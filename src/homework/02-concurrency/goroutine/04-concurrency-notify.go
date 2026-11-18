// 题目：给定订单号切片 []string{"O1", "O2", "O3", "O4"}，每个订单在协程里 Sleep 50 毫秒后打印「处理完成: 订单号」；
// 用 WaitGroup 等全部处理完再打印「收银台下班」。

package main

import (
	"fmt"
	"sync"
	"time"
)

func notifyExample(orders []string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(50 * time.Millisecond)

	for _, order := range orders {
		fmt.Println("处理完成:", order)
	}
}

func ConcurrencyNotify() {
	var wg sync.WaitGroup

	orders := []string{
		"order-1",
		"order-2",
		"order-3",
		"order-4",
	}

	for _, order := range orders {
		wg.Add(1)
		go notifyExample([]string{order}, &wg)
	}

	wg.Wait()

	fmt.Println("收银台下班")
}
