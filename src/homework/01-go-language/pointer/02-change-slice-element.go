// 题目：写 firstToX(s []string)，把第一个元素改成 "X"。传入 []string{"A", "B", "C"}，调用后打印切片，确认元素已被改变。

package main

import "fmt"

func firstToX(s []string) {
	s[0] = "X"
}

func ChangeSliceElement() {
	s := []string{"A", "B", "C"}

	firstToX(s)

	fmt.Println(s) // [X B C]：切片描述符按值传，但底层数组是共享的
}
