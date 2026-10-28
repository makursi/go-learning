// 题目：用 var 显式声明一个 string 和一个 int 变量，再用短声明同时定义两个变量（注意名字与值要对上），最后依次打印。

package main

import "fmt"

func DeclareVariables() {
	var person string = "小满"
	var num int = 18

	name, age := "choria", 19

	fmt.Println(name, age)
	fmt.Println(person)
	fmt.Println(num)
}
