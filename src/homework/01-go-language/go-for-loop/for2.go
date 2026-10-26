package main

import "fmt"

// break / continue：遍历 1 到 20，遇到能被 3 整除的数就 continue 跳过；
// 一旦遇到大于 15 的数就 break 结束。观察最终打印了哪些数。

func For2() {
	for i := 1; i <= 20; i++ {
		if i > 15 {
			break
		}

		if i%3 == 0 {
			continue
		}

		fmt.Println(i)
	}
}
