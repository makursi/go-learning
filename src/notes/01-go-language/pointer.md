# pointer

## 值传递与包含引用的类型

Go 中**所有函数参数都是按值传递（pass by value）**的。调用函数时，会复制参数的值并传递给函数。

### 值的表示（Representation of values）

来自语言规范《Properties of types and values → Representation of values》部分：

- 预声明类型（以及数组和结构体）的值是自包含的：每个这样的值都包含其全部数据的完整副本，对应类型的变量存储整个值。例如，数组变量为数组的所有元素提供存储空间。这些类型的零值是类型特定的，永远不是 `nil`。

- 非 nil 的指针、函数、切片、映射和通道值**包含对底层数据的引用**，这些底层数据可能被多个值共享：
  - 指针值是对持有指针基类型值的变量的引用。
  - 函数值包含对（可能是匿名的）函数和封闭变量的引用。
  - **切片值包含切片的长度、容量，以及对它底层数组的引用**。
  - **映射或通道值是对映射或通道的实现特定数据结构的引用**。

接口值可能是自包含的，也可能包含对底层数据的引用，取决于接口的动态类型。预声明标识符 `nil` 是那些值可以包含引用的类型的零值。

当多个值共享底层数据时，修改一个值可能会改变另一个。例如，修改切片的一个元素会改变所有共享该数组的切片中底层数组的对应元素。

### 赋值时的行为

来自语言规范《Assignment statements》相关说明：

当一个值被赋给变量时，只有存储在变量中的数据被替换。如果该值包含引用，赋值会**复制引用，但不会复制被引用的数据**（例如切片的底层数组）。

官方示例（来自规范）：

```go
var s1 = []int{1, 2, 3}
var s2 = s1                    // s2 存储 s1 的切片描述符
s1 = s1[:1]                    // s1 的长度变为 1，但它仍与 s2 共享底层数组
s2[0] = 42                     // 设置 s2[0] 也会改变 s1[0]
fmt.Println(s1, s2)            // 打印 [42] [42 2 3]
```

映射同理：

```go
var m1 = make(map[string]int)
var m2 = m1                    // m2 存储 m1 的映射描述符
```

### 切片与映射的行为（来自 Effective Go）

**切片**：

切片的运行时数据结构（包含指针、长度和容量）是按值传递的。虽然函数可以修改切片的元素（因为元素位于共享的底层数组中），但切片本身（描述符）是按值复制的。因此，如果函数需要改变切片的长度或容量（例如通过 `append`），通常需要返回新的切片。

官方说明（Effective Go）：

> We must return the slice afterwards because, although `Append` can modify the elements of `slice`, the slice itself (the run-time data structure holding the pointer, length, and capacity) is passed by value.

**映射**：

> Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.

简单示例（基于官方描述的行为）：

```go
package main

import "fmt"

func modifySlice(s []int) {
    s[0] = 100  // 修改会影响调用者，因为共享底层数组
}

func modifyMap(m map[string]int) {
    m["key"] = 42  // 修改会影响调用者，因为共享底层数据结构
}

func main() {
    s := []int{1, 2, 3}
    modifySlice(s)
    fmt.Println(s)  // [100 2 3]

    m := map[string]int{"key": 1}
    modifyMap(m)
    fmt.Println(m)  // map[key:42]
}
```
