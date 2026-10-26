//数值转换：声明一个 float64 变量，值为 9.99，用 int() 转成整数并打印；再把结果用 float64() 转回去，观察小数部分是否还在。

package main

func IntCovers() {

	floatNum := 9.99
	intNum := int(floatNum)
	println("Float:", floatNum)
	println("Int:", intNum)
	convertedBack := float64(intNum)
	println("Converted Back:", convertedBack)
}
