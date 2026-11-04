# 协程（Goroutine）

**Go 官方文档中关于协程（goroutine）的内容**


### 什么是 goroutine

*goroutine* 是由 Go 运行时管理的轻量级线程。


“go” 语句会启动一个函数调用的执行，使其作为一个独立的并发控制线程（即 *goroutine*）在同一地址空间中运行。

### 启动 goroutine

使用 `go` 语句启动一个新的 goroutine：

```
go f(x, y, z)
```

这会启动一个新的 goroutine 来运行：

```
f(x, y, z)
```

`f`、`x`、`y` 和 `z` 的求值发生在当前 goroutine 中，而 `f` 的执行发生在新的 goroutine 中。


规范中进一步说明：

函数值和参数会像往常一样在调用方的 goroutine 中求值，但与普通调用不同，程序执行不会等待被调用的函数完成。相反，该函数会在一个新的 goroutine 中独立开始执行。当函数终止时，其 goroutine 也会终止。如果函数有任何返回值，这些返回值在函数完成时会被丢弃。



规范中的示例：

```
go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true }} (c)
```

### 官方简单代码示例

以下是 A Tour of Go 中的示例：

```
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
```

### 关于共享内存与同步

goroutine 运行在同一地址空间中，因此对共享内存的访问必须进行同步。`sync` 包提供了有用的原语，不过在 Go 中你通常不太需要它们，因为还有其他原语可用。


### 程序退出时的行为

当该函数调用返回时，程序退出。它不会等待其他（非 `main`）goroutine 完成。


### 内存模型中的相关说明

启动新 goroutine 的 `go` 语句会在该 goroutine 开始执行之前完成同步。


