//结构体值 vs 指针：定义 Counter struct { Value int }。
// 分别写 addByValue(c Counter) 和 addByPointer(c *Counter)，都让 Value + 1。
// 在 main 里各调用一次，对比哪种方式能让外面的 Counter 真的加 1。

package main

type Counter struct {
	Value int
}

func addByValue(c Counter) {
	c.Value += 1
}

func addByPointer(c *Counter) {
	c.Value += 1 // 指针接收者的字段访问会自动解引用，不必写 (*c).Value
}
