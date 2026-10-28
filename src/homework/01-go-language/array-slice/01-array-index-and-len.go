// 题目：声明一个长度为 5 的整型数组，打印下标 2 的元素和整个数组的长度。

package main

import "fmt"

func ArrayIndexAndLen() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array[2])
	fmt.Println(len(array))
}
