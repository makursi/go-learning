package main

import "fmt"

type Book struct {
	Title string
	Price float64
}

func Struct() {
	//结构体：定义一个 Book 结构体，包含 Title（书名）和 Price（价格）两个字段，创建一本书的实例并打印书名和价格。
	book := Book{
		Title: "Go Programming",
		Price: 29.99,
	}
	fmt.Printf("Book: %s, Price: %.2f\n", book.Title, book.Price)
}
