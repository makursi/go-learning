package main

import "fmt"

//切片截取：给定 s := []int{10, 20, 30, 40, 50}，分别打印 s[1:4]、s[:3]、s[2:]、s[:]，对照文中表格核对结果。

func Slice1() {
	s := []int{10, 20, 30, 40, 50}
	fmt.Println(s[1:4]) //20,30,40
	fmt.Println(s[:3])  //10,20,30
	fmt.Println(s[2:])  //30,40,50
	fmt.Println(s[:])   //10,20,30,40,50
}
