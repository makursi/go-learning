package main

import "fmt"

func SliceArray() {
	// 数组与切片：
	// 定义一个长度为 3 的整型数组，赋值为 [10, 20, 30] 并打印；
	// 定义一个空切片，用 append 依次追加 1, 2, 3，打印切片内容和长度 len。

	array := [3]int{10, 20, 30}
	fmt.Println(array)

	slice := []int{}
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	fmt.Println(slice)
	fmt.Println(len(slice))
}
