# 互斥锁与读写锁（Mutex 与 RWMutex）

**Go 互斥锁与读写锁学习笔记**  

### 概述
`sync` 包提供了基本的同步原语，包括互斥锁。除了 `Once` 和 `WaitGroup` 外，大多数类型主要用于低级库例程。更高层次的同步更适合通过 channel 和通信来完成。包含本包中定义的类型的值在首次使用后不应被复制。

### Mutex（互斥锁）
`Mutex` 是一个互斥锁。  
`Mutex` 的零值是一个未锁定的互斥锁。  

`Mutex` 在首次使用后不得被复制。

在 Go 内存模型的术语中，对 `Mutex.Unlock` 的第 n 次调用“同步先于”对 `Mutex.Lock` 的第 m 次调用（对于任意 n < m）。成功的 `TryLock` 调用等价于一次 `Lock` 调用。失败的 `TryLock` 调用不建立任何“同步先于”关系。

**主要方法：**

- `func (m *Mutex) Lock()`  
  锁定 m。如果锁已被使用，调用的 goroutine 会阻塞，直到互斥锁可用。

- `func (m *Mutex) Unlock()`  
  解锁 m。如果在进入 `Unlock` 时 m 未锁定，则会发生运行时错误。  
  锁定的 `Mutex` 不与特定的 goroutine 关联。允许一个 goroutine 锁定 `Mutex`，然后安排另一个 goroutine 解锁它。

- `func (m *Mutex) TryLock() bool`（Go 1.18 添加）  
  尝试锁定 m 并报告是否成功。  
  注意：虽然 `TryLock` 的正确用法确实存在，但它们很少见，使用 `TryLock` 通常是互斥锁使用中存在更深层次问题的信号。

**官方示例（来自 A Tour of Go）：**

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```

说明：通过在 `Inc` 方法中用 `Lock` 和 `Unlock` 包围代码块，以及在 `Value` 方法中使用 `defer` 确保解锁，保证了只有一个 goroutine 可以同时访问共享的 map。

### RWMutex（读写互斥锁）
`RWMutex` 是一个读者/写者互斥锁。  
该锁可以由任意数量的读者或单个写者持有。  
`RWMutex` 的零值是一个未锁定的互斥锁。  

`RWMutex` 在首次使用后不得被复制。

如果任何 goroutine 在锁已被一个或多个读者持有时调用 `RWMutex.Lock`，则并发的 `RWMutex.RLock` 调用将阻塞，直到写者获取（并释放）锁，以确保锁最终对写者可用。  
注意：这禁止了递归读锁定。  
`RWMutex.RLock` 不能升级为 `RWMutex.Lock`，`RWMutex.Lock` 也不能降级为 `RWMutex.RLock`。

在 Go 内存模型的术语中，对 `RWMutex.Unlock` 的第 n 次调用“同步先于”对 `Lock` 的第 m 次调用（对于任意 n < m），与 `Mutex` 相同。  
对于任何 `RLock` 调用，都存在一个 n，使得对 `Unlock` 的第 n 次调用“同步先于”该 `RLock` 调用，并且对应的 `RWMutex.RUnlock` 调用“同步先于”对 `Lock` 的第 n+1 次调用。

**主要方法：**

- `func (rw *RWMutex) Lock()`  
  为写入锁定 rw。如果锁已经为读或写而锁定，则 `Lock` 会阻塞，直到锁可用。

- `func (rw *RWMutex) Unlock()`  
  为写入解锁 rw。如果在进入 `Unlock` 时 rw 未为写入而锁定，则会发生运行时错误。  
  与 `Mutex` 一样，锁定的 `RWMutex` 不与特定的 goroutine 关联。一个 goroutine 可以 `RLock`（或 `Lock`）一个 `RWMutex`，然后安排另一个 goroutine 来 `RUnlock`（或 `Unlock`）它。

- `func (rw *RWMutex) RLock()`  
  为读取锁定 rw。  
  它不应用于递归读锁定；被阻塞的 `Lock` 调用会排除新读者获取锁。详见 `RWMutex` 类型文档。

- `func (rw *RWMutex) RUnlock()`  
  撤销一次 `RWMutex.RLock` 调用；它不影响其他同时进行的读者。  
  如果在进入 `RUnlock` 时 rw 未为读取而锁定，则会发生运行时错误。

- `func (rw *RWMutex) RLocker() Locker`  
  返回一个实现了 `Locker` 接口的对象，其 `Lock` 和 `Unlock` 方法通过调用 `rw.RLock` 和 `rw.RUnlock` 实现。

- `func (rw *RWMutex) TryLock() bool`（Go 1.18 添加）  
  尝试为写入锁定 rw 并报告是否成功。  
  注意：虽然 `TryLock` 的正确用法确实存在，但它们很少见，使用 `TryLock` 通常是互斥锁使用中存在更深层次问题的信号。

- `func (rw *RWMutex) TryRLock() bool`（Go 1.18 添加）  
  尝试为读取锁定 rw 并报告是否成功。  
  注意：虽然 `TryRLock` 的正确用法确实存在，但它们很少见，使用 `TryRLock` 通常是互斥锁使用中存在更深层次问题的信号。

**简单使用示意（严格依据官方方法描述）：**

```go
var rw sync.RWMutex
var data int

// 写操作
rw.Lock()
data = 42
rw.Unlock()

// 读操作
rw.RLock()
_ = data
rw.RUnlock()
```

多个读者可以同时持有读锁，但写锁是独占的。当写者等待时，新的读者会被阻塞，以保证写者最终能获得锁。

### 补充说明（来自官方文档）
- `Locker` 接口定义为：
  ```go
  type Locker interface {
      Lock()
      Unlock()
  }
  ```
  `Mutex` 和 `RWMutex`（通过 `RLocker`）都实现了该接口。

- 文档反复强调：这些类型在首次使用后不得被复制。

