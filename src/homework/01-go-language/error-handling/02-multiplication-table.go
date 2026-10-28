// 题目：写一个函数，传入 1～9 打印对应的乘法表；小于 1 或大于 9 时返回 error。调用失败后打印一句「其余逻辑」，确认它仍会执行。

package main

import (
	"errors"
	"fmt"
)

func printTimesTable(n int) error {
	if n > 9 || n < 1 {
		return errors.New("请输入1~9的整数")
	}

	fmt.Println("以下是其余逻辑:")

	for i := 1; i <= n; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d * %d = %d  ", k, i, i*k)
		}
		fmt.Println()
	}

	return nil
}

func MultiplicationTable() {
	if err := printTimesTable(5); err != nil {
		fmt.Println("调用失败", err)
	}

	fmt.Println("其余逻辑")
}
