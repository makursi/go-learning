# type-conversion

**Go 类型转换**

Go 是强类型语言，不同类型之间通常不能隐式转换，需要显式转换。数字类型之间的转换由语言本身支持，而与字符串、布尔值相关的转换则主要通过标准库 `strconv` 包实现。

## 数字的转换

根据语言规范，非常量数字值之间的转换遵循以下规则：

在整数类型之间转换时，如果源值是有符号整数，则进行符号扩展到隐式无限精度；否则进行零扩展。然后截断以适应目标类型的大小。

在整数与浮点类型之间转换时，结果会向零舍入到最接近的可表示的浮点值。如果超出范围，则结果是对应的无穷大。

在浮点类型之间转换时，结果会舍入到目标类型的精度。

常量转换有更严格的规则：常量必须能在目标类型中精确表示（或按规则舍入）。

简单示例：

```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = float64(i)  // 整数转浮点
    var u uint = uint(f)        // 浮点转无符号整数

    fmt.Println(i, f, u)
}
```

## 数字转字符串

标准库 `strconv` 包提供将数字转换为字符串的函数。

最常用的是 `Itoa`（int 转字符串）和 `FormatInt`、`FormatUint`、`FormatFloat`。

官方文档示例：

```go
s := strconv.Itoa(-42)
s := strconv.FormatInt(-42, 16)
s := strconv.FormatUint(42, 16)
s := strconv.FormatFloat(3.1415, 'E', -1, 64)
```

简单完整示例：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    i := 10
    s := strconv.Itoa(i)
    fmt.Printf("%T, %v\n", s, s)  // string, 10

    f := 3.1415
    sf := strconv.FormatFloat(f, 'f', 2, 64)
    fmt.Println(sf)  // 3.14
}
```

`Itoa` 等价于 `FormatInt(int64(i), 10)`。

## 字符串转数字

`strconv` 包提供 `Atoi`、`ParseInt`、`ParseUint`、`ParseFloat` 等函数将字符串解析为数字。

官方文档说明：最常见的是 `Atoi`（字符串转 int）和 `ParseInt`、`ParseFloat` 等。

```go
i, err := strconv.Atoi("-42")
f, err := strconv.ParseFloat("3.1415", 64)
i, err := strconv.ParseInt("-42", 10, 64)
u, err := strconv.ParseUint("42", 10, 64)
```

解析函数返回最宽的类型（`float64`、`int64`、`uint64`）。如果指定了较窄的位宽，结果可再转换为该较窄类型且不会丢失数据。

简单示例：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "2147483647" // 最大的 int32
    i64, err := strconv.ParseInt(s, 10, 32)
    if err != nil {
        fmt.Println(err)
        return
    }
    i := int32(i64)
    fmt.Println(i)

    // 常用 Atoi
    n, err := strconv.Atoi("-42")
    if err == nil {
        fmt.Println(n)
    }
}
```

解析失败时返回错误（如 `strconv.ErrSyntax` 或 `strconv.ErrRange`）。

## 字符串转布尔值

使用 `strconv.ParseBool`。

官方文档：`ParseBool` 返回字符串所表示的布尔值。它接受 `1`、`t`、`T`、`TRUE`、`true`、`True`、`0`、`f`、`F`、`FALSE`、`false`、`False`。其他任何值都会返回错误。

简单示例：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    v := "true"
    if b, err := strconv.ParseBool(v); err == nil {
        fmt.Printf("%T, %v\n", b, b)  // bool, true
    }

    // 其他合法值
    b1, _ := strconv.ParseBool("1")
    b2, _ := strconv.ParseBool("FALSE")
    fmt.Println(b1, b2)  // true false
}
```

## 布尔值转字符串

使用 `strconv.FormatBool`。

官方文档：`FormatBool` 根据布尔值返回 `"true"` 或 `"false"`。

简单示例：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := strconv.FormatBool(true)
    fmt.Println(s)  // true

    s = strconv.FormatBool(false)
    fmt.Println(s)  // false
}
```

也可使用 `AppendBool` 将结果追加到字节切片。
