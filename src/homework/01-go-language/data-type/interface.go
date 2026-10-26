package main

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Book2 struct {
	Title  string
	Author string
	Pages  int
}

func Interface() {
	//接口：定义一个 Shape 接口，包含 Area() float64 方法；再定义 Rectangle（矩形）结构体并实现该接口，创建实例后打印面积。
	//  综合练习：将上述 Book 结构体放入切片中，存储至少 3 本书，遍历切片打印每本书的信息。
	shape := Rectangle{width: 5, height: 10}
	bookslice := []Book2{
		{Title: "Book 1", Author: "Author 1", Pages: 100},
		{Title: "Book 2", Author: "Author 2", Pages: 200},
		{Title: "Book 3", Author: "Author 3", Pages: 300},
	}
	println("Rectangle Area:", shape.Area())
	for _, book := range bookslice {
		println("Book:", book.Title, "Author:", book.Author, "Pages:", book.Pages)
	}
}
