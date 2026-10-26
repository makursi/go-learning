package main

import "fmt"

//普通 for：声明一个 1 到 10 的循环，打印所有偶数。

func For() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
