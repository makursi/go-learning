# 通道（Channel）

**Go 通道（Channel）**

### 通道类型

通道提供了一种机制，让并发执行的函数通过发送和接收指定元素类型的值进行通信。未初始化通道的值为 `nil`。

通道类型的语法为：

```
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .
```

可选的 `<-` 操作符指定通道的方向（发送或接收）。如果给出了方向，通道就是单向的；否则是双向的。通道可通过赋值或显式转换被约束为仅发送或仅接收。

```
chan T          // 可用于发送和接收类型为 T 的值
chan<- float64  // 只能用于发送 float64
<-chan int      // 只能用于接收 int
```

两个通道类型相同，当且仅当它们的元素类型相同且方向相同。

### 创建通道

使用内置函数 `make` 创建通道：

```
make(T)          // 无缓冲通道
make(T, n)       // 缓冲通道，缓冲区大小为 n
```

例如：

```
ci := make(chan int)            // 整数无缓冲通道
cj := make(chan int, 0)         // 整数无缓冲通道
cs := make(chan *os.File, 100)  // 指向 File 的指针的缓冲通道
```

容量（以元素数量计）设置通道缓冲区的大小。如果容量为零或不存在，通道是无缓冲的，通信只有在发送方和接收方都准备好时才会成功。否则，通道是缓冲的，如果缓冲区未满（发送）或非空（接收），通信可在不阻塞的情况下成功。`nil` 通道永远不会准备好进行通信。

### 发送语句

发送语句在通道上发送一个值。通道表达式必须是通道类型，通道方向必须允许发送操作，且要发送的值的类型必须可赋值给通道的元素类型。

```
SendStmt = Channel "<-" Expression .
Channel  = Expression .
```

通道和值表达式在通信开始前都会被求值。通信会阻塞，直到发送可以继续进行。无缓冲通道上的发送可以在接收方准备好时进行。缓冲通道上的发送可以在缓冲区有空间时进行。向已关闭的通道发送会导致运行时 panic。向 `nil` 通道发送会永久阻塞。

```
ch <- 3  // 向通道 ch 发送值 3
```

### 接收操作符

对于通道类型的操作数 `ch`，接收操作 `<-ch` 的值是从通道 `ch` 接收到的值。通道方向必须允许接收操作，接收操作的类型是通道的元素类型。表达式会阻塞，直到有值可用。从 `nil` 通道接收会永久阻塞。对已关闭通道的接收操作总是可以立即进行，在接收完之前发送的值后，会生成元素类型的零值。

```
v1 := <-ch
v2 = <-ch
f(<-ch)
<-strobe // 等待时钟脉冲并丢弃接收到的值
```

多值赋值形式的接收操作会报告接收到的值是否是在通道关闭前发送的：

```
x, ok = <-ch
x, ok := <-ch
var x, ok = <-ch
var x, ok T = <-ch
```

`ok` 的值为 `true` 表示该值是由成功的发送操作传递的，为 `false` 表示因为通道已关闭且为空而生成了零值。

### 关闭通道

内置函数 `close` 用于关闭通道：

```
close(ch)
```

`close(ch)` 记录不再有更多值会发送到该通道。如果 `ch` 是只接收通道，则是错误。向已关闭的通道发送或关闭已关闭的通道会导致运行时 panic。关闭 `nil` 通道也会导致运行时 panic。调用 `close` 后，在接收完之前发送的值后，接收操作会返回通道类型的零值且不阻塞。多值接收操作会返回接收到的值以及通道是否已关闭的指示。

### 通道的其他特性

单个通道可被任意数量的 goroutine 用于发送语句、接收操作以及内置函数 `cap` 和 `len`，无需进一步同步。通道充当先进先出队列。例如，如果一个 goroutine 在通道上发送值，第二个 goroutine 接收它们，则值按发送顺序接收。

无缓冲通道将通信（交换值）与同步（保证两个计算（goroutine）处于已知状态）结合起来。接收方总是阻塞，直到有数据可接收。如果通道是无缓冲的，发送方会阻塞，直到接收方已接收该值。如果通道有缓冲区，发送方仅在值被复制到缓冲区之前阻塞；如果缓冲区已满，则意味着等待某个接收方检索一个值。

通道是一等值，可以像任何其他值一样分配和传递。

### 简单示例（来自 Effective Go）

启动一个排序并等待其完成：

```
c := make(chan int)  // 分配一个通道。
// 在 goroutine 中启动排序；完成后在通道上发出信号。
go func() {
    list.Sort()
    c <- 1  // 发送信号；值无关紧要。
}()
doSomethingForAWhile()
<-c   // 等待排序完成；丢弃发送的值。
```

使用缓冲通道作为信号量限制并发：

```
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // 等待活动队列排空。
    process(r)  // 可能需要很长时间。
    <-sem       // 完成；为下一个启用“槽”。
}
```

### range 遍历通道

`range` 子句可以遍历通道上陆续发送的值。规范里的语法：

```
RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
```

规范对通道的规定只有一句：**产生的迭代值是通道上连续发送的值，直到通道被关闭**——通道是 `nil` 时，range 表达式永久阻塞。

> For channels, the iteration values produced are the successive values sent on the channel until the channel is closed. If the channel is `nil`, the range expression blocks forever.

这句话解释了 `range` 为什么「不给你 `ok`」：循环终止的唯一条件是「通道已关闭且已排空」。所以**关闭是发送方的责任**，发送方不 `close`，接收方的 `range` 就永远等下去。另外，`range` 对通道只允许一个迭代变量（数组、切片、map 才允许两个）。

`for v := range ch` 与手写循环等价：

```
for {
    v, ok := <-ch
    if !ok {
        break
    }
    // 使用 v
}
```

例子：

```go
func main() {
	ch := make(chan string, 2)
	ch <- "go"
	ch <- "home"
	close(ch) // 少了这一句，下面的 range 会永久阻塞

	for v := range ch {
		fmt.Println(v)
	}
}
```

输出：

```
go
home
```

### select 多路复用

`select` 从一组发送、接收操作中挑一个继续执行。语法（规范）：

```
SelectStmt = "select" "{" { CommClause } "}" .
CommClause = CommCase ":" StatementList .
CommCase   = "case" ( SendStmt | RecvStmt ) | "default" .
```

执行过程规范分了几步，要点三条：

1. 进入 `select` 时，所有 `case` 里的通道操作数按源码顺序各求值一次；接收语句左侧的赋值目标此时**不**求值。
2. 有一个以上通信可以继续，就通过均匀伪随机选择挑一个；否则若有 `default` 就执行它；否则阻塞，直到至少一个通信可以继续。
3. `nil` 通道上的通信永远不会继续，因此只含 `nil` 通道又没有 `default` 的 `select` 永久阻塞。

> If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement blocks until at least one of the communications can proceed.

由这三条推出来的常用写法：

- **同时就绪是随机的**：两个通道都有值可读时，别指望写在前面先执行、也别指望谁先发送，只能指望两者最终都会被执行。
- **`default` 让 select 变成非阻塞**：立即返回，用来做「有就读、没有就跳过」。

```go
func main() {
	ch := make(chan int, 1)

	select {
	case v := <-ch:
		fmt.Println("读到", v)
	default:
		fmt.Println("通道为空，不等待")
	}

	ch <- 7

	select {
	case v := <-ch:
		fmt.Println("读到", v)
	default:
		fmt.Println("通道为空，不等待")
	}
}
```

输出：

```
通道为空，不等待
读到 7
```

- **置 `nil` 摘掉 case**：通道关闭后把它设成 `nil`，这个 `case` 就永远不就绪，这是「循环消费多个通道直到全部关完」的收尾写法；不置 `nil` 的话，关闭的通道每轮都返回零值，循环退不出来。

```go
for chA != nil || chB != nil {
	select {
	case v, ok := <-chA:
		if !ok {
			chA = nil // 摘掉，否则每轮都读到零值
			continue
		}
		// 使用 v
	}
}
```

- **超时**：`case <-time.After(d)` 在 `d` 之后让这个 `case` 就绪，和结果通道赛跑就实现了超时；注意 `time.After` 每次调用都新建一个定时器，循环里高频使用要换成 `time.NewTimer` 并复用。

```go
func wait(ch <-chan string, timeout time.Duration) (string, bool) {
	select {
	case v := <-ch:
		return v, true
	case <-time.After(timeout):
		return "", false
	}
}
```

### 常见注意点

- 关闭是发送方的动作：对只接收通道 `close` 是编译错误，重复 `close` 与关闭 `nil` 通道都是运行时 panic。
- 关闭之后接收不会阻塞：先收完缓冲区里剩下的值，再返回元素类型的零值且 `ok` 为 `false`；关闭之后继续发送则 panic。
- `range` 与 `for` + `select` 的退出都依赖 `close`（或额外的停止通道），发送方漏了关闭，接收方就永久阻塞。
- 多个 `case` 同时就绪时选哪个是随机的，不要依赖书写顺序。

