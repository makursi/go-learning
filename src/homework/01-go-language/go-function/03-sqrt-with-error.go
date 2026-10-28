// 题目：写一个 sqrt(n float64) (float64, error)：n < 0 时返回错误，否则返回平方根；分别用正数和负数调用，正确处理 err。

package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("输入数字不能为负数")
	}

	return math.Sqrt(n), nil
}

func SqrtWithError() {
	result, err := sqrt(16)
	if err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("平方根：", result)
	}

	result, err = sqrt(-4)
	if err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("平方根：", result)
	}
}
