package main

import "fmt"

func String() {
	//字符串：声明两个字符串变量，分别存储你的姓名和学校名称，用 + 拼接后打印，中间用换行符 \n 分隔。
	name := "Golang"
	school := "XYZ University"
	fmt.Println(name + "\n" + school)
}
