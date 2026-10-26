package main

import (
	"fmt"
	"strconv"
)

//布尔转换：把字符串 "1"、"false"、"yes" 分别传给 strconv.ParseBool()，打印每次转换的结果和错误（这次不要用 _ 忽略错误，看看哪些能成功）。

func BoolToString() {
	s1 := "1"
	s2 := "false"
	s3 := "yes"

	s1bool, err := strconv.ParseBool(s1)
	fmt.Println(s1bool, err)

	s2bool, err := strconv.ParseBool(s2)
	fmt.Println(s2bool, err)

	s3bool, err := strconv.ParseBool(s3)
	fmt.Println(s3bool, err)

}
