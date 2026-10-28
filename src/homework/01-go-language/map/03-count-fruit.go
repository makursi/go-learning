// 题目：用 map[string]int 统计切片 []string{"苹果", "香蕉", "苹果", "橘子", "香蕉", "苹果"} 里每种水果出现的次数，并遍历打印结果。

package main

import "fmt"

func CountFruit() {
	fruits := []string{"苹果", "香蕉", "苹果", "橘子", "香蕉", "苹果"}

	count := make(map[string]int)

	// range 一个切片时每轮给两个值：索引与元素本身。这里不要索引，用 _ 舍掉。
	for _, value := range fruits {
		// 读取不存在的键会得到值类型的零值（这里是 int 的 0），所以可以直接自增。
		count[value]++
	}

	fmt.Println(count)

	// range 一个 map 时每轮给键与值，两个都留着。
	for fruit, num := range count {
		fmt.Println(fruit, "出现了", num, "次")
	}
}
