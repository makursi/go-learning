//panic 与 recover：写一个函数 mustPositive(n int)，n <= 0 时 panic。
// 在 main 里用 defer + recover 捕获，打印 panic 信息；并在调用后面再写一句 fmt.Println("其余逻辑")，验证它不会执行。

package main

func mustPositive(n int) (int, error) {

	if n <= 0 {
		panic("错误")
	}

	return 1, nil
}
