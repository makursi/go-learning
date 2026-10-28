// 题目：用「只写条件」的 for（不写初始化段和步进段）计算 1 + 2 + … + 100 的和并打印。

package main

import "fmt"

func WhileSumOneToHundred() {
	i := 1
	sum := 0

	for i <= 100 {
		sum += i
		i++
	}

	fmt.Println(sum)
}
