# go-for-loop

以下内容全部依据 Go 官方文档（语言规范 [The Go Programming Language Specification](https://go.dev/ref/spec)、[A Tour of Go](https://go.dev/tour/)、[Effective Go](https://go.dev/doc/effective_go)）整理。

## 普通 for 循环

Go 只有一种循环结构：`for` 循环。

基本形式有三个由分号分隔的组成部分：
- 初始化语句（init statement）：在第一次迭代前执行
- 条件表达式（condition expression）：每次迭代前求值
- 后置语句（post statement）：每次迭代结束后执行

初始化语句通常是短变量声明，在其中声明的变量仅在 `for` 语句作用域内可见。当布尔条件求值为 `false` 时，循环停止。

与 C、Java 或 JavaScript 不同，三个组成部分外没有括号，花括号 `{ }` 始终是必需的。

```go
package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
```

`for` 循环还可写成类似 while 的形式，或无限循环：

```go
// 类似 C 的 while
for condition { }

// 类似 C 的 for(;;)
for { }
```

示例（条件形式）：

```go
for a < b {
    a *= 2
}
```

## range 循环

带 `range` 子句的 `for` 语句会遍历数组、切片、字符串或 map 的所有条目，从 channel 接收的值，以及（Go 1.22 起）从 0 到上限的整数，或（Go 1.23 起）传递给迭代器函数 yield 的值。对每个条目，它将迭代值赋给对应的迭代变量（如果存在），然后执行循环体。

```go
RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
```

### 遍历 map

对 map 使用 `range` 时，第一个值是键，第二个值（可选）是对应的值。

```go
for k, v := range myMap {
    log.Printf("key=%v, value=%v", k, v)
}
```

只使用一个变量时，得到的是键：

```go
for k := range myMap {
    // 使用 k
}
```

### 遍历字符串

对字符串使用 `range` 时，第一个值是索引（字节位置），第二个值是对应的 Unicode 码点（rune）。`range` 按 Unicode 码点迭代，而不是按字节。

```go
for i, r := range s {
    // i 是字节索引，r 是 rune
}
```

只使用一个变量时，得到的是索引。

## break 与 continue

### break：跳出循环

`break` 语句终止同一函数内最内层的 `for`、`switch` 或 `select` 语句的执行。

```go
BreakStmt = "break" [ Label ] .
```

如果带有标签，该标签必须是外层 `for`、`switch` 或 `select` 语句的标签，执行会终止该语句。

```go
OuterLoop:
    for i = 0; i < n; i++ {
        for j = 0; j < m; j++ {
            switch a[i][j] {
            case nil:
                state = Error
                break OuterLoop
            case item:
                state = Found
                break OuterLoop
            }
        }
    }
```

### continue：跳过本次

`continue` 语句开始最内层封闭 `for` 循环的下一次迭代，通过将控制转移到循环块的末尾。该 `for` 循环必须在同一函数内。

```go
ContinueStmt = "continue" [ Label ] .
```

如果带有标签，该标签必须是外层 `for` 语句的标签，执行会推进该循环。

```go
RowLoop:
    for y, row := range rows {
        for x, data := range row {
            if data == endOfRow {
                continue RowLoop
            }
            row[x] = data + bias(x, y)
        }
    }
```

### 带标签的 break

`break` 和 `continue` 都可带可选标签，用于指定要跳出或继续的外层循环（或 switch/select）。标签放在语句前，并以冒号结尾。

上面 `OuterLoop` 与 `RowLoop` 的示例即展示了带标签的用法。

## goto

`goto` 语句将控制转移到同一函数内带有对应标签的语句。

```go
GotoStmt = "goto" Label .
```

示例：

```go
goto Error
```

执行 `goto` 语句时，不得使任何在 `goto` 点尚未进入作用域的变量进入作用域。

标签语句本身可作为 `goto`、`break` 或 `continue` 的目标：

```go
LabeledStmt = Label ":" Statement .
Label       = identifier .
```