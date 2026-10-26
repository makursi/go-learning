package main

import "fmt"

// 前置与中间插入：以 []int{1, 2, 3, 4, 5} 为起点，先把 -1, 0 插到最前面；再在元素 3 后面插入 100, 200，最终打印完整切片。
func Slice2() {
	slice := []int{1, 2, 3, 4, 5}
	slice = append([]int{-1, 0}, slice...)
	fmt.Println(slice)
	slice = append(slice[:5], append([]int{100, 200}, slice[5:]...)...)
	//append([]int{100, 200}, slice[5:]...)  [4,5]前插入[100,200]
	//append(slice[:5], append([]int{100, 200}, slice[5:]...)...) [-1,0,1,2,3] 后插入[100,200,4,5]
	fmt.Println(slice)
}
