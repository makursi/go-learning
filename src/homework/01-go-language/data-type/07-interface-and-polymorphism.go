// 题目：定义 Shape 接口（Area() float64）与 Rectangle 结构体并实现该接口，创建实例后打印面积；
// 综合练习：把书目结构体放进切片，存至少 3 本，遍历打印每本的信息。

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

// Book2 与第一题的 Book 同包重名，所以叫 Book2。
type Book2 struct {
	Title  string
	Author string
	Pages  int
}

func InterfaceAndPolymorphism() {
	shape := Rectangle{width: 5, height: 10}

	books := []Book2{
		{Title: "Book 1", Author: "Author 1", Pages: 100},
		{Title: "Book 2", Author: "Author 2", Pages: 200},
		{Title: "Book 3", Author: "Author 3", Pages: 300},
	}

	println("Rectangle Area:", shape.Area())

	for _, book := range books {
		println("Book:", book.Title, "Author:", book.Author, "Pages:", book.Pages)
	}
}
