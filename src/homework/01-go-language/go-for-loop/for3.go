package main

import "fmt"

// 类 while：用「只写条件」的 for（不要初始化段和步进段）
// 计算 1 + 2 + ... + 100 的和并打印。

func For3() {
	i := 1
	sum := 0

	for i <= 100 {
		sum += i
		i++
	}

	fmt.Println(sum)
}
