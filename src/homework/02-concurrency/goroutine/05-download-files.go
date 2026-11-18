// 题目：定义 download(file string, wg *sync.WaitGroup)，内部 Sleep 100 毫秒后打印「已下载: 文件名」；
// 用循环 + go 启动 []string{"a.txt", "b.txt", "c.txt"} 的下载（注意循环变量拷贝或按参数传入），最后 Wait 并打印「全部下载完成」。

package main

import (
	"fmt"
	"sync"
	"time"
)

func download(file string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("已下载:", file)
}

func DownloadFiles() {
	var wg sync.WaitGroup

	files := []string{"a.txt", "b.txt", "c.txt"}

	for _, file := range files {
		wg.Add(1)
		go download(file, &wg) // 按参数传入，避免捕获循环变量
	}

	wg.Wait()

	fmt.Println("全部下载完成")
}
