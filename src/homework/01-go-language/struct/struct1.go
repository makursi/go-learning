package main

import "fmt"

// 嵌套：定义 Address（城市、街道）和 User（姓名、地址），创建用户并打印「姓名住在城市的街道」。
// 嵌入：把上一题的 Address 改成嵌入到 User，用提升后的字段直接打印城市
type Address struct {
	City   string
	Street string
}

type User struct {
	Name string
	Address
}

func Struct1() {
	user1 := User{
		Name: "choria",
		Address: Address{
			City:   "datong",
			Street: "pingcheng",
		},
	}

	fmt.Println(user1)
	fmt.Println(user1.City)
}
