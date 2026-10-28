// 题目：定义 Book 结构体（书名、作者、页数），用至少两种方式创建实例并打印；
// 再给 Book 写值接收者方法 Info() 打印信息、指针接收者方法 AddPages(n int) 增加页数，调用后验证页数有变。

package main

import "fmt"

type Book struct {
	Name   string
	Author string
	Pages  int
}

// Info 是值接收者：拿到的是副本，只读不写。
func (b Book) Info() {
	fmt.Println(b)
}

// AddPages 是指针接收者：改的是调用者那个 Book 本身。
func (b *Book) AddPages(n int) {
	b.Pages += n
}

func DefineAndMethods() {
	book := Book{
		Name:   "go program",
		Author: "Alice",
		Pages:  10,
	}

	book1 := Book{
		Name: "none",
	}

	fmt.Println(book)
	fmt.Println(book1)

	book1.Info()
	book1.AddPages(1)
	fmt.Println(book1.Pages)
}
