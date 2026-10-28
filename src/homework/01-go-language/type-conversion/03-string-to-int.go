// 题目：声明两个字符串 "15" 和 "27"，分别用 strconv.Atoi() 转成整数后相加，打印和。

package main

import (
	"fmt"
	"strconv"
)

func StringToInt() {
	s1 := "15"
	s2 := "27"

	i1, _ := strconv.Atoi(s1)
	i2, _ := strconv.Atoi(s2)

	sum := i1 + i2
	fmt.Println(sum)
}
