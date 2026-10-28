// 题目：定义 Book 结构体（Title、Price 两个字段），创建一本书的实例并打印书名和价格。

package main

import "fmt"

type Book struct {
	Title string
	Price float64
}

func StructBasics() {
	book := Book{
		Title: "Go Programming",
		Price: 29.99,
	}

	fmt.Printf("Book: %s, Price: %.2f\n", book.Title, book.Price)
}
