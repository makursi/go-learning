package main

import "fmt"

func main() {
	fruits := []string{
		"苹果", "香蕉", "苹果",
		"橘子", "香蕉", "苹果",
	}

	fruitCount := make(map[string]int)

	// map 的零值是 0，所以直接自增即可，不用先判断键是否存在
	for _, fruit := range fruits {
		fruitCount[fruit]++
	}

	// 注意：range map 的顺序随机，输出顺序每次可能不同
	for fruit, count := range fruitCount {
		fmt.Println(fruit, count)
	}
}
