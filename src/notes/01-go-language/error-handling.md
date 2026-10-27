# error-handling

## 错误是值（Errors are values）

Go 把错误当作普通值来处理，而不是异常。函数在可能失败时，通常会把 `error` 作为最后一个返回值返回。调用方通过检查该返回值是否为 `nil` 来判断是否发生了错误。这种设计鼓励在错误发生的地方显式检查，而不是依赖隐式的异常抛出与捕获。

官方强调：始终检查错误。错误是值，因此可以用 Go 语言的全部能力来处理它们。

## error 接口

`error` 是内置接口类型，声明如下：

```go
type error interface {
    Error() string
}
```

任何实现了 `Error() string` 方法的类型都可以作为错误值使用。`nil` 的 `error` 表示没有错误。

## 创建错误

标准库 `errors` 包提供最基本的创建方式：

```go
import "errors"

err := errors.New("empty name")
```

也可以用 `fmt.Errorf`（更常用，可格式化）：

```go
import "fmt"

err := fmt.Errorf("math: square root of negative number %g", f)
```

官方建议：错误字符串通常不以大写字母开头，因为打印时常常会被前缀拼接，例如：

```go
err := TryInverseTachyonPulse()
if err != nil {
    fmt.Printf("failed to solve problem: %s\n", err)
}
```

## 返回与处理错误的基本模式

函数通常返回 `(结果, error)`。成功时返回 `nil` 错误，失败时返回非 `nil` 错误。

官方教程示例（greetings 包）：

```go
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }
    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
```

调用方处理：

```go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := greetings.Hello("")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)
}
```

运行后会输出类似：

```
greetings: empty name
exit status 1
```

这是 Go 中最常见的错误处理方式：把错误作为值返回，让调用方检查。

Effective Go 也展示了典型风格（成功路径保持在第一缩进层级）：

```go
f, err := os.Open(name)
if err != nil {
    return err
}
// 使用 f
```

## 自定义错误类型

当需要携带额外信息时，可以定义自己的类型并实现 `Error()` 方法。

官方示例：

```go
type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
    return fmt.Sprintf("math: square root of negative number %g", float64(f))
}
```

另一个官方示例（json 包风格）：

```go
type SyntaxError struct {
    msg    string // description of error
    Offset int64  // error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string { return e.msg }
```

调用方可用类型断言或类型 switch 检查：

```go
if err := dec.Decode(&val); err != nil {
    if serr, ok := err.(*json.SyntaxError); ok {
        line, col := findLine(f, serr.Offset)
        return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
    }
    return err
}
```

Wiki 建议命名约定：错误类型以 `Error` 结尾，错误变量以 `Err` 或 `err` 开头。

## 哨兵错误（Sentinel errors）

对于调用方需要特别处理的已知错误条件，可定义包级变量：

```go
var ErrNotFound = errors.New("not found")
```

比较时使用 `errors.Is`（见下文），而不是直接 `==`（在包装场景下更安全）。

## 错误包装（Wrapping，Go 1.13 起）

从 Go 1.13 开始，官方为错误包装提供了明确支持。一个错误可以“包装”另一个错误，形成错误链（或树）。

约定是实现 `Unwrap() error`（或 `Unwrap() []error`）方法。

最简单的包装方式是使用 `fmt.Errorf` 的 `%w` 动词：

```go
if err != nil {
    return fmt.Errorf("decompress %v: %w", name, err)
}
```

`%w` 与 `%v` 在格式化文本上相同，但会让返回的错误实现 `Unwrap`，从而暴露底层错误。

是否包装需要权衡：
- 包装会把底层错误暴露给调用方，使其成为 API 的一部分。
- 如果不想暴露实现细节，应使用 `%v` 或重新创建错误，而不是 `%w`。

## 检查包装后的错误：errors.Is 与 errors.As

`errors.Is(err, target)` 检查错误树中是否包含与 `target` 匹配的错误（优先于直接 `==`）：

```go
if errors.Is(err, fs.ErrNotExist) {
    fmt.Println("file does not exist")
}
```

`errors.As(err, &target)` 检查错误树中是否存在可赋给 `target` 类型的错误，并赋值：

```go
var pathError *fs.PathError
if errors.As(err, &pathError) {
    fmt.Println("Failed at path:", pathError.Path)
}
```

较新版本还提供了泛型形式 `errors.AsType[E error](err error) (E, bool)`，使用更直接。

也可以自定义 `Is` 或 `As` 方法来控制匹配行为。

`errors.Unwrap(err)` 返回一层包装，但通常更推荐直接使用 `Is` 或 `As`。

### 合并多个错误：errors.Join（Go 1.20 起）

```go
err := errors.Join(err1, err2, err3)
```

`Join` 会丢弃其中的 `nil`，如果全部为 `nil` 则返回 `nil`。结果错误的文本是各个错误字符串以换行拼接。

## 减少重复错误检查的技巧

官方博客指出，可以通过编程错误值来减少样板代码。例如 `bufio.Scanner` 把错误收集起来，只在最后通过 `Err()` 报告；或者自定义类似 `errWriter` 的类型，把多次写入的错误集中处理。

核心思想仍然是：错误是值，可以用语言本身简化处理，但最终必须检查错误。
