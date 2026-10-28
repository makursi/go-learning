// 题目：创建一个 map[string]int，存 3 位同学的姓名与分数，打印其中一位的分数；
// 再添加一位新同学、删除其中一位，最后遍历打印所有键值对。

package main

import "fmt"

func MapBasics() {
	students := map[string]int{
		"Alice":   85,
		"Bob":     95,
		"Charlie": 90,
	}

	fmt.Println("Bob's score:", students["Bob"])

	students["David"] = 88
	delete(students, "Alice")

	fmt.Println("All students and their scores:")
	for name, score := range students {
		fmt.Printf("%s: %d\n", name, score)
	}
}
