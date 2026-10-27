//业务错误对比：改写文中的 nn，允许传入 1～9 任意整数打印对应乘法表；小于 1 或大于 9 时返回 error。调用失败后打印一句「其余逻辑」，确认它仍会执行。

package main

import (
	"errors"
	"fmt"
)

func nn(a int) (bool, error) {
	if a > 9 || a < 1 {
		return false, errors.New("请输入1~9的整数")
	}
	fmt.Println("以下是其余逻辑:")

	for i := 1; i <= a; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d * %d = %d  ", k, i, i*k)
		}
		fmt.Println()
	}
	return true, nil
}
