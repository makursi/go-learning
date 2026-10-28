// 题目：以 []int{1, 2, 3, 4, 5} 为起点，先把 -1、0 插到最前面，再在元素 3 后面插入 100、200，最后打印完整切片。

package main

import "fmt"

func InsertFrontAndMiddle() {
	slice := []int{1, 2, 3, 4, 5}

	slice = append([]int{-1, 0}, slice...)
	fmt.Println(slice)

	// append([]int{100, 200}, slice[5:]...) 先把 [4 5] 接到 [100 200] 后面，
	// 再整段接在 slice[:5]（也就是 [-1 0 1 2 3]）后面。
	slice = append(slice[:5], append([]int{100, 200}, slice[5:]...)...)
	fmt.Println(slice)
}
