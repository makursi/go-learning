//遍历统计：用 map[string]int 统计字符串切片 []string{"苹果", "香蕉", "苹果", "橘子", "香蕉", "苹果"} 中每种水果出现的次数，并遍历打印结果。

package main

import "fmt"

func Map2() {
	fruits := []string{"苹果", "香蕉", "苹果", "橘子", "香蕉", "苹果"}

	// 1. 创建一个空的 map
	count := make(map[string]int)

	// 2. 遍历切片，统计水果出现的次数
	//这里的 range fruits 数组和切片时会在每次循环中提供两个值：
	// - 第一个值：元素的索引。
	// - 第二个值：元素本身。
	//不需要索引，舍去
	for _, value := range fruits {

		//由于 "苹果" 还没有作为键存在于 map 中，读取 count["苹果"] 会得到 int 的零值，也就是 0。
		//利用了Go 的一个重要特性：读取 map 中不存在的键时，会返回该值类型的零值。
		count[value]++

	}
	fmt.Println(count)
	// 3. 遍历 map，打印统计结果
	//range 关于 map 返回两个迭代值
	// 	key：Map 中的键。
	// value：这个键对应的值。

	for fruit, num := range count {
		fmt.Println(fruit, "出现了", num, "次")
	}
}
