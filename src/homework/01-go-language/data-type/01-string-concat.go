// 题目：声明两个字符串变量（姓名、学校），用 + 拼接后打印，中间用换行符 \n 分隔。

package main

import "fmt"

func StringConcat() {
	name := "Golang"
	school := "XYZ University"

	fmt.Println(name + "\n" + school)
}
