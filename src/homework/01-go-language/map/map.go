package main

import "fmt"

// 声明与读写：创建一个 map[string]string，存自己的姓名、城市、爱好；修改城市，再新增一个「职业」键，打印整个 map。
func Map() {
	person := map[string]string{
		"Name":  "choria",
		"City":  "datong",
		"Hobby": "program",
	}

	person["Career"] = "student"

	fmt.Println(person)
}
