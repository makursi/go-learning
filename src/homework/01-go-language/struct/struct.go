package main

import "fmt"

// 定义与初始化：定义 Book 结构体（书名、作者、页数），用至少两种方式创建实例并打印。
// 方法：给 Book 写值接收者方法 Info() 打印信息；再写指针接收者方法 AddPages(n int) 增加页数，调用后验证页数有变。
type Book struct {
	Name   string
	Author string
	Pages  int
}

func (b Book) Info() {
	fmt.Println(b)
}

func (b *Book) AddPages(n int) {
	b.Pages += n
}
func Struct() {

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
	//使用Info打印信息，使用AddPages增加页数并打印
	book1.Info()
	book1.AddPages(1)
	fmt.Println(book1.Pages)
}
