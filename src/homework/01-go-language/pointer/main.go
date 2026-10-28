package main

import "fmt"

func main() {

	//创建setZero函数并修改值
	num := 42
	fmt.Println(num)
	setZero(num)
	fmt.Println(num)

	//切片修改元素
	s := []string{
		"A", "B", "C",
	}
	firstToX(s)

	fmt.Println(s)

	//map 增删改
	scores := map[string]int{
		"语文": 90,
		"数学": 85,
		"英语": 92,
	}

	updateScore(scores, "英语", 30) // 改

	fmt.Println(scores)

	updateScore(scores, "物理", 10) // 增：科目不存在时等价于 0 + delta

	fmt.Println(scores)

	removeSubject(scores, "语文") // 删

	fmt.Println(scores)

	//结构体的值引用和指针引用

	count := Counter{
		Value: 1,
	}

	addByPointer(&count)
	fmt.Println(count)
	addByValue(count)
	fmt.Println(count)

}
