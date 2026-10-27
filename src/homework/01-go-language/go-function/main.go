package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// 基本函数：写一个 max(a, b int) int，返回较大的那个数；在 main 里调用并打印结果。

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 多返回值：写一个 swap(a, b int) (int, int)，交换两个整数并返回；用 x, y := swap(1, 2) 验证。

func swap(a, b int) (int, int) {

	var temp int
	temp = a
	a = b
	b = temp
	return a, b
}

// 错误处理：写一个 sqrt(n float64) (float64, error)：n < 0 时返回错误（可用 errors.New），否则返回平方根（可用 math.Sqrt）。
// 在 main 里分别用正数和负数调用，正确处理 err。

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("输入数字不能为负数")
	}

	sqrtNum := math.Sqrt(n)
	return sqrtNum, nil
}

// 匿名函数：在 main 里用匿名函数实现「打印 1 到 5」，立刻执行；再把同一个逻辑赋给变量后调用一次。

// 综合练习：写一个 parseAndDouble(s string) (int, error)，用 strconv.Atoi 把字符串转成整数后乘以 2；转换失败则把错误返回。在 main 里测试 "21" 和 "abc" 两种输入。

func parseAndDouble(s string) (int, error) {
	num, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	}

	return num * 2, nil //返回函数的默认空值nil
}
func main() {

	// 最大值函数
	maxNum := max(10, 20)
	fmt.Println(maxNum)
	//交换函数
	fmt.Println(swap(10, 20))
	// 错误处理函数
	// 正数调用
	result, err := sqrt(16)
	if err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("平方根：", result)
	}
	// 负数调用
	result, err = sqrt(-4)
	if err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("平方根：", result)
	}

	//匿名函数

	func() {
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
		}
	}()

	//综合练习

	convNum1, err := parseAndDouble("21")
	convNum2, err := parseAndDouble("abc")

	if err != nil {
		fmt.Println("发生错误：", err)
	} else {
		fmt.Println("计算结果：", convNum1)
	}
	if err != nil {
		fmt.Println("发生错误：", err)
	} else {
		fmt.Println("计算结果：", convNum2)
	}
}
