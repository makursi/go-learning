package main

import "fmt"

// defer 顺序：在 main 里注册三个 defer，打印 A、B、C，再打印一句「函数体」，观察输出顺序是否为「函数体 → C → B → A」。
func main() {

	// 返回 error
	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("计算错误", err)
	} else {
		fmt.Println("计算结果为:", result)
	}

	// 业务错误对比

	_, err = nn(5)

	if err != nil {
		fmt.Println("调用失败", err)
	}

	fmt.Println("其余逻辑")

	// defer 先进后出
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
	fmt.Println("函数体")

	// defet + recover 的使用

	defer func() {
		if info := recover(); info != nil {
			fmt.Println("捕获到程序错误:", info)
		}
	}()
	result, err = mustPositive(-1)

	if err != nil {
		fmt.Println("执行失败")
	} else {
		fmt.Println("执行成功", result)
	}

	fmt.Println("其余逻辑")
}
