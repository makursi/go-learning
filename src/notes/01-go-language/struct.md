# struct

## 定义结构体

根据 Go 语言规范，结构体（struct）是由一系列命名元素（称为字段）组成的序列，每个字段都有名称和类型。字段名称可以显式指定，也可以通过嵌入字段隐式指定。在同一个结构体中，非空白字段名称必须唯一。

结构体类型的语法如下：

```
StructType = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl  = (IdentifierList Type | EmbeddedField) [ Tag ] .
```

简单示例（来自官方规范与 A Tour of Go）：

```go
type Vertex struct {
    X, Y int
}

v := Vertex{1, 2}          // 结构体字面量
v2 := Vertex{X: 1}         // 仅指定部分字段，其余为零值
v3 := Vertex{}             // 所有字段为零值
p := &Vertex{1, 2}         // 指向结构体的指针
```

空结构体也是合法的：

```go
struct {}
```

结构体字段可以通过点号访问。当通过结构体指针访问字段时，Go 允许省略显式解引用，直接使用 `p.X` 的形式。

## 结构体嵌套

结构体字段的类型可以是另一个结构体类型，从而形成嵌套结构。嵌套结构体字段需要通过完整的字段路径访问。

根据语言规范中的类型定义与复合字面量规则，嵌套是普通字段声明的自然结果。示例如下（基于官方规范中关于结构体字段类型的描述）：

```go
type Point3D struct {
    x, y, z float64
}

type Line struct {
    p, q Point3D
}

origin := Point3D{}
line := Line{origin, Point3D{y: -4, z: 12.3}}
```

访问嵌套字段时使用点号链：

```go
line.p.x
line.q.z
```

结构体类型 `T` 不能包含类型为 `T` 的字段，也不能包含以数组或结构体形式直接或间接包含 `T` 的类型（指针、切片、函数等除外）。

## 结构体嵌入

根据 Go 语言规范，声明时只给出类型而没有显式字段名的字段称为**嵌入字段**（embedded field）。

嵌入字段必须是类型名 `T`，或指向非接口类型名的指针 `*T`，且 `T` 本身不能是指针类型或类型参数。未限定的类型名即作为字段名。

示例（直接来自语言规范）：

```go
// 包含四个嵌入字段的结构体
struct {
    T1        // 字段名为 T1
    *T2       // 字段名为 T2
    P.T3      // 字段名为 T3
    *P.T4     // 字段名为 T4
    x, y int  // 普通字段
}
```

嵌入字段的字段和方法可以**提升**（promoted）。如果 `x.f` 是合法选择器，且 `f` 来自嵌入字段，则称 `f` 被提升。提升后的字段可像普通字段一样使用。

方法集规则（来自规范）：

- 若结构体 `S` 包含嵌入字段 `T`，则 `S` 和 `*S` 的方法集都包含接收者为 `T` 的提升方法；`*S` 的方法集还包含接收者为 `*T` 的提升方法。
- 若结构体 `S` 包含嵌入字段 `*T`，则 `S` 和 `*S` 的方法集都包含接收者为 `T` 或 `*T` 的提升方法。

## 方法

Go 没有类，但可以为类型定义方法。方法是带有特殊接收者参数的函数。接收者出现在 `func` 关键字与方法名之间。

### 值接受

接收者类型为值类型 `T` 时，方法操作的是接收者的副本。

示例（来自 A Tour of Go）：

```go
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

v := Vertex{3, 4}
fmt.Println(v.Abs())
```

### 指针接受

接收者类型为指针类型 `*T` 时，方法可以修改指向的值，并且通常更高效（避免大结构体的复制）。

示例（来自 A Tour of Go）：

```go
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

v := Vertex{3, 4}
v.Scale(10)        // 即使 v 是值，也可调用指针接收者方法
```

语言规范与 Tour 均指出：带指针接收者的方法在调用时，接收者可以是值或指针（编译器自动取地址）；带值接收者的方法在调用时，接收者可以是值或指针（编译器自动解引用）。一般建议同一类型的所有方法统一使用值接收者或指针接收者。

## 匿名结构体

结构体类型可以不通过 `type` 声明命名，直接以类型字面量形式使用，称为匿名结构体（unnamed struct type）。

根据语言规范和官方演讲材料（Stupid Gopher Tricks），匿名结构体常用于一次性数据、模板数据或 JSON 处理。

示例：

```go
// 变量声明中的匿名结构体
var t struct {
    i int
    s string
}

// 复合字面量中的匿名结构体
data := struct {
    Title               string
    Firstname, Lastname string
    Rank                int
}{
    "Dr", "Carl", "Sagan", 7,
}

// 嵌套匿名结构体（常用于 JSON）
var data struct {
    ID     int
    Person struct {
        Name string
        Job  string
    }
}
```

空的匿名结构体 `struct{}` 也是常见惯用法。

## 泛型结构体

自 Go 1.18 起，类型可以带有类型参数，从而支持泛型结构体。

根据 A Tour of Go 与语言规范中的类型参数规则，泛型结构体示例如下：

```go
// 表示可保存任意类型值的单链表
type List[T any] struct {
    next  *List[T]
    value T
}
```

类型参数出现在类型名后的方括号中。实例化时提供具体类型参数即可使用该泛型结构体。类型参数本身不能作为嵌入字段的类型名。