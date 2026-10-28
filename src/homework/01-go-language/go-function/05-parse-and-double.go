// 题目：写一个 parseAndDouble(s string) (int, error)，用 strconv.Atoi 把字符串转成整数后乘以 2；转换失败则返回错误。分别用 "21" 和 "abc" 测试。

package main

import (
	"fmt"
	"strconv"
)

func parseAndDouble(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return num * 2, nil // nil 表示没有错误
}

func ParseAndDouble() {
	// 每次调用后立刻判错：err 只有一个，晚判就会读到下一次调用的结果。
	convNum1, err := parseAndDouble("21")
	if err != nil {
		fmt.Println("发生错误：", err)
	} else {
		fmt.Println("计算结果：", convNum1)
	}

	convNum2, err := parseAndDouble("abc")
	if err != nil {
		fmt.Println("发生错误：", err)
	} else {
		fmt.Println("计算结果：", convNum2)
	}
}
