//数字转字符串：声明一个 int 变量表示年龄，用 strconv.Itoa() 转成字符串，再用 fmt.Printf 打印类型和值（确认类型是 string）。

package main

import (
	"fmt"
	"strconv"
)

func IntToString() {

	var age int = 10
	agetostring := strconv.Itoa(age)

	fmt.Printf("类型: %T, 值: %v\n", agetostring, agetostring)
}
