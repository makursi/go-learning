# map

## 语法

根据 Go 语言规范，map 类型的语法为：

```
MapType = "map" "[" KeyType "]" ElementType .
KeyType = Type .
```

示例：

```go
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

键类型必须是可比较的类型（comparison operators `==` 和 `!=` 必须完全定义）。因此键类型不能是函数、map 或 slice。值类型可以是任意类型，包括另一个 map。

## 定义与初始化

map 是无序的元素集合，元素类型相同，通过唯一的键进行索引。未初始化的 map 值为 `nil`。

声明一个 map 变量：

```go
var m map[string]int
```

map 是引用类型（类似指针或切片），上述 `m` 的值为 `nil`，它不指向已初始化的 map。

使用内置函数 `make` 初始化 map：

```go
m = make(map[string]int)
```

`make` 可接受可选的容量提示：

```go
make(map[string]int, 100)
```

初始容量不限制其大小，map 会根据存储的元素数量增长（`nil` map 除外）。

也可以使用 map 字面量初始化：

```go
commits := map[string]int{
    "rsc": 3711,
    "r":   2138,
    "gri": 1908,
    "adg": 912,
}
```

初始化空 map 的字面量形式与 `make` 功能相同：

```go
m = map[string]int{}
```

来自 A Tour of Go 的示例：

```go
type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}
```

若顶层类型只是类型名，可在字面量元素中省略该类型名。

## 读写元素

插入或更新元素：

```go
m[key] = elem
```

检索元素：

```go
elem = m[key]
```

如果请求的键不存在，则返回值类型的零值。

示例：

```go
m := make(map[string]int)
m["route"] = 66
i := m["route"]   // i == 66
j := m["root"]    // j == 0（零值）
```

## 判断键是否存在

使用双值赋值形式测试键是否存在：

```go
elem, ok = m[key]
// 或短声明
elem, ok := m[key]
```

如果键存在，`ok` 为 `true`；如果不存在，`ok` 为 `false`，且 `elem` 为零值。

仅测试存在性而不需要值时：

```go
_, ok := m["route"]
```

示例（来自 A Tour of Go）：

```go
m := make(map[string]int)
m["Answer"] = 42
v, ok := m["Answer"]
fmt.Println("The value:", v, "Present?", ok)
```

## 删除元素

使用内置函数 `delete` 删除指定键的元素：

```go
delete(m, key)
```

`delete` 不返回任何值。如果 map 为 `nil` 或键不存在，则 `delete` 是空操作（no-op）。

示例：

```go
delete(m, "Answer")
```

从 Go 1.21 起，还可使用内置函数 `clear` 删除 map 中的所有条目，使 map 变为空（`len(m) == 0`）：

```go
clear(m)
```

## 获取长度

内置函数 `len` 返回 map 中已定义键的数量：

```go
n := len(m)
```

对于 map 类型 `map[K]T`，`len` 返回定义的键的数量，且在执行过程中可能发生变化。

## 遍历 map

使用 `range` 关键字遍历 map：

```go
for key, value := range m {
    fmt.Println("Key:", key, "Value:", value)
}
```

遍历时，第一个值为键，第二个值为对应的元素值。

根据语言规范，对 map 进行 `range` 时：

- 第一个迭代值为键（类型为键类型）
- 第二个迭代值为对应的 map 元素（类型为元素类型）

迭代顺序未指定，且不保证每次迭代相同。若需要稳定顺序，必须维护单独的数据结构来指定顺序。

## 作为函数参数

map 是引用类型。将 map 作为参数传递或赋值时，复制的是引用（map 描述符），而不会复制底层数据。

语言规范中的示例：

```go
var m1 = make(map[string]int)
var m2 = m1                    // m2 存储 m1 的 map 描述符
m1["foo"] = 42                 // 设置 m1["foo"] 也会改变 m2["foo"]
fmt.Println(m2["foo"])         // 打印 42
```

因此，函数内部对 map 的修改会反映到调用者传入的原始 map 上。

## 常见注意点

- **nil map 行为**：`nil` map 在读取时表现得像空 map，但向 `nil` map 写入会导致运行时 panic。`nil` map 等价于空 map，只是不能添加元素。
- **键类型限制**：键必须是可比较类型。可比较类型包括布尔、数值、字符串、指针、通道和接口类型，以及仅包含这些类型的结构体或数组。slice、map 和函数不能作为键。
- **零值便利性**：检索不存在的键会返回零值，这在某些场景下很有用（例如用 `map[T]bool` 实现类似集合的结构）。
- **并发安全**：map 不是并发安全的。同时读写 map 的行为未定义。如果需要从多个并发执行的 goroutine 读写 map，必须使用同步机制（例如 `sync.RWMutex`）进行保护。
- **迭代顺序**：`range` 遍历 map 的顺序不确定，且不保证稳定。需要有序遍历时，应单独维护排序后的键切片。
- **容量提示**：`make(map[K]V, n)` 创建的初始空间大约可容纳 `n` 个元素，但具体行为依赖实现，map 仍会按需增长。