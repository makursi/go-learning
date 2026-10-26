package main

import "fmt"

//range 练习：定义切片 []string{"春", "夏", "秋", "冬"}，分别用「只要下标」「只要值」「下标+值」三种 range 写法打印。

func For1() {
	season := []string{"春", "夏", "秋", "冬"}

	for idx, _ := range season {
		fmt.Println(idx)
	}

	for _, value := range season {
		fmt.Println(value)
	}

	for idx, value := range season {
		fmt.Println(idx, value)
	}
}
