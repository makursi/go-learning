//切片声明与追加：声明切片 []string{"唱", "跳", "rap"}，用 append 在尾部追加 "篮球"，打印切片内容、len 和 cap。

package main

import "fmt"

func Slice() {
	kun := []string{
		"唱", "跳", "rap",
	}

	kun = append(kun, "篮球")

	fmt.Println(kun)
	fmt.Println(len(kun))
	fmt.Println(cap(kun))
}
