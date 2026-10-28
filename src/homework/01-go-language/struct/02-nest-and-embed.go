// 题目：定义 Address（城市、街道）与 User（姓名、地址），创建用户并打印；
// 再把 Address 改成嵌入字段，用提升后的字段直接取城市。

package main

import "fmt"

type Address struct {
	City   string
	Street string
}

type User struct {
	Name    string
	Address // 嵌入字段：没有字段名，类型名就是字段名
}

func NestAndEmbed() {
	user1 := User{
		Name: "choria",
		Address: Address{
			City:   "datong",
			Street: "pingcheng",
		},
	}

	fmt.Println(user1)
	fmt.Println(user1.City) // 提升上来的字段，等价于 user1.Address.City
}
