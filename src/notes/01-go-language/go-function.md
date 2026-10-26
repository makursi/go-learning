# go-function

## 函数定义

在 Go 中，函数使用 `func` 关键字声明。函数可以有零个或多个参数，以及零个或多个返回值。参数和返回值都需要指定类型。

简单示例（来自官方语言规范与 Effective Go 的惯例）：

```go
func add(x int, y int) int {
    return x + y
}
```

连续相同类型的参数可以省略前面的类型声明：

```go
func add(x, y int) int {
    return x + y
}
```

### 多返回值

Go 函数可以返回多个结果。这是语言的核心特性之一，在标准库中非常常见（例如返回计算结果与错误）。

示例（来自官方文档与 Effective Go 风格）：

```go
func swap(x, y string) (string, string) {
    return y, x
}
```

调用时：

```go
a, b := swap("hello", "world")
```

也可以为返回值命名（命名返回值）。命名返回值在函数入口处会被初始化为零值，并且可以在 `return` 时直接使用裸 `return`：

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

### 注意事项

- 返回值的顺序与声明顺序一致，调用方必须按相同顺序接收。
- 如果不需要某个返回值，可使用空白标识符 `_` 丢弃。
- 命名返回值主要用于文档说明和简化某些清理逻辑，但应谨慎使用，避免可读性下降。
- 所有返回值在函数开始时都会被初始化为零值。

## 错误处理

Go 中错误是值。惯例是函数返回一个或多个结果，最后一个返回值通常是 `error` 类型。`nil` 表示没有错误。

示例（来自官方教程 “Return and handle an error”）：

```go
func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
```

调用方应立即检查错误：

```go
message, err := Hello("")
if err != nil {
    // 处理错误
    return
}
// 使用 message
```

错误是接口类型，只有一个方法：

```go
type error interface {
    Error() string
}
```

创建错误可使用 `errors.New` 或 `fmt.Errorf`。

## 匿名函数

Go 支持函数字面量（function literal），即没有名字的函数。它们是一等公民，可作为值赋值、传递或返回。

示例（来自 A Tour of Go）：

```go
func main() {
    // 赋值给变量
    f := func(x, y int) int {
        return x + y
    }
    fmt.Println(f(1, 2))
}
```

立即执行的匿名函数：

```go
func() {
    fmt.Println("hello")
}()
```

匿名函数可以形成闭包：它引用并绑定了外部作用域的变量。

示例（来自 A Tour of Go 的 Function closures）：

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
```

每次调用 `adder` 都会创建独立的 `sum` 变量，返回的闭包与该变量绑定。

### 注意事项

- 闭包捕获的是变量本身（按引用），而不是当时的值。在循环中创建闭包时尤其需要注意这一点。
- 匿名函数可以访问并修改其词法作用域中的变量。
- 函数字面量本身是表达式，可以出现在任何需要值的地方。

## 泛型函数

自 Go 1.18 起，函数可以声明类型参数，从而编写可适用于多种类型的通用代码。

类型参数写在方括号中，并带有约束（constraint）。

示例（来自官方教程 “Getting started with generics”）：

```go
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

调用时可以显式提供类型实参，也可以让编译器推断：

```go
SumIntsOrFloats[string, int64](ints)
SumIntsOrFloats(ints) // 类型推断
```

也可以把约束提取为接口：

```go
type Number interface {
    int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

## defer 关键字

`defer` 语句会把一个函数调用推入栈中，该调用会在外层函数返回时执行（无论是正常返回还是 panic）。

示例（来自 A Tour of Go）：

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
// 输出：
// hello
// world
```

### 多个 defer：后进先出（LIFO）

多个 `defer` 按照后进先出的顺序执行。

官方博客示例：

```go
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
// 输出：3210
```

### 实用案例：保证资源释放

`defer` 最常见的用途是保证资源（文件、锁、连接等）被释放，即使函数有多个返回点。

官方博客示例：

```go
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```

### 再记两点

1. **参数在 defer 语句执行时求值**（而不是在真正调用时）。
   ```go
   func a() {
       i := 0
       defer fmt.Println(i) // 打印 0
       i++
   }
   ```

2. **defer 函数可以读写命名返回值**。
   ```go
   func c() (i int) {
       defer func() { i++ }()
       return 1
   }
   // 返回 2
   ```

## 用接口模拟「面向对象」能力

Go 没有类和继承。它通过**方法**和**接口**提供面向对象风格的能力。

方法是带有接收者（receiver）的函数：

```go
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

接口定义一组方法签名。任何实现了这些方法的类型都隐式满足该接口（无需显式声明 “implements”）：

```go
type Abser interface {
    Abs() float64
}
```

只要 `Vertex` 实现了 `Abs` 方法，它就可以被当作 `Abser` 使用。这实现了多态，而无需继承层次。

Go 强调组合而非继承，接口用于描述行为，从而实现灵活的抽象。