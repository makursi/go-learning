//切片改元素：写 firstToX(s []string)，把第一个元素改成 "X"。在 main 里传入 []string{"A", "B", "C"}，调用后打印切片，确认已改变。

package main

func firstToX(s []string) {
	s[0] = "X"
}
