// 题目：写一个 1 到 10 的 for 循环，打印其中所有偶数。

package main

import "fmt"

func PrintEvenNumbers() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
