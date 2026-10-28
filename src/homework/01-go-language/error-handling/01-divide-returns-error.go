// 题目：写一个 divide(a, b int) (int, error)，b 为 0 时返回错误，否则返回商；分别在 b = 2 与 b = 0 时调用并处理。

package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}

	return a / b, nil
}

func DivideReturnsError() {
	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("计算错误", err)
	} else {
		fmt.Println("计算结果为:", result)
	}

	// 除数为零：错误分支要真的走一次
	result, err = divide(10, 0)

	if err != nil {
		fmt.Println("计算错误", err)
	} else {
		fmt.Println("计算结果为:", result)
	}
}
