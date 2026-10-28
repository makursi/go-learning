// 题目：声明一个 float64 变量值为 9.99，用 int() 转成整数并打印；再把结果用 float64() 转回去，观察小数部分是否还在。

package main

import "fmt"

func FloatIntRoundTrip() {
	floatNum := 9.99
	intNum := int(floatNum)
	fmt.Println("Float:", floatNum)
	fmt.Println("Int:", intNum)

	convertedBack := float64(intNum)
	fmt.Println("Converted Back:", convertedBack)
}
