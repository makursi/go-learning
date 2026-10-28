// 题目：定义 Counter struct { Value int }，分别写 addByValue(c Counter) 与 addByPointer(c *Counter) 各让 Value 加一。
// 各调用一次，对比哪种方式能让外面的 Counter 真的加一。

package main

import "fmt"

type Counter struct {
	Value int
}

func addByValue(c Counter) {
	c.Value += 1
}

func addByPointer(c *Counter) {
	c.Value += 1 // 指针接收者的字段访问会自动解引用，不必写 (*c).Value
}

func StructValueVsPointer() {
	count := Counter{Value: 1}

	addByPointer(&count)
	fmt.Println(count) // {2}：真的改了

	addByValue(count)
	fmt.Println(count) // {2}：改的是副本
}
