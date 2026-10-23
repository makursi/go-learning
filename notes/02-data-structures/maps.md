# Map

> 讲义：[lectures/02-data-structures/maps](../../lectures/02-data-structures/maps)　作业：[homework/02-data-structures/fruit-count](../../homework/02-data-structures/fruit-count)

Map（哈希表，hash table）是键值对集合，键必须可比较（`comparable`）—— 切片、map、函数不能做键。

## 零值

nil map 可以读、可以 `len`、可以 `range`，但**写入会 panic**：

```go
var m map[string]int
fmt.Println(m["x"], len(m)) // 0 0
m["x"] = 1                  // panic: assignment to entry in nil map
```

所以声明之后必须 `make` 或字面量初始化再写。

## 初始化

```go
a := make(map[string]int)              // 空 map，可写
b := map[string]int{"语文": 90}        // 字面量
c := make(map[string]int, 16)          // 预估容量，减少扩容
```

`make` 的第二个参数是**提示容量**，不是上限：超出照样扩容。

## 增删改查

| 操作 | 写法 | 备注 |
| --- | --- | --- |
| 增 / 改 | `m[k] = v` | 键不存在即新增，存在即覆盖 |
| 查 | `v := m[k]` | 缺失时得到**零值**，无法区分「是零」和「没有」 |
| 查（区分缺失） | `v, ok := m[k]` | `ok` 为 `false` 表示键不存在 |
| 删 | `delete(m, k)` | 键不存在也不报错 |
| 长度 | `len(m)` | |

「缺失返回零值」是 Go map 最容易写出 bug 的地方：`if m[k] == 0` 无法判断键是否真的存在。

## 遍历

`range` 只给键，`range` 两个变量才给键值：

```go
for k := range m { }
for k, v := range m { }
```

**遍历顺序是随机的**，每次运行都可能不同。不要依赖顺序，需要有序输出就先收集键再 `sort`。

## 并发

map 不是并发安全的。多 goroutine 同时写会直接 panic（`concurrent map writes`）。要并发就配 `sync.Mutex` 或 `sync.RWMutex`，或按场景换 `sync.Map`。

## 常见场景

- 计数：`count[k]++`（零值是 0，所以不用先判断存在）
- 去重 / 集合：`seen[k] = struct{}{}`
- 缓存
- 分组：`map[string][]T`

---

*这是一篇示例笔记，用来定下格式，可以删。*
