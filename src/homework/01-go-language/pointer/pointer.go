//值传递：写 setZero(n int)，在函数内把 n 设为 0。在 main 里传入 42，调用前后各打印一次，确认外面仍是 42。

package main

func setZero(n int) {
	n = 0
}
