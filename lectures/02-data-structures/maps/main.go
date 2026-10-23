package main

import "fmt"

func main() {
	// 字面量初始化
	scores := map[string]int{
		"语文": 90,
		"数学": 95,
	}

	// 增、改：键不存在即新增，存在即覆盖
	scores["英语"] = 88
	scores["数学"] = 97

	// 查：comma-ok 区分「缺失」与「值为零」
	if v, ok := scores["英语"]; ok {
		fmt.Println("英语:", v)
	}
	if _, ok := scores["物理"]; !ok {
		fmt.Println("物理: 缺考")
	}

	// 删：键不存在也不报错
	delete(scores, "语文")

	// 遍历：顺序随机，每次运行都可能不同
	for subject, score := range scores {
		fmt.Printf("%s=%d\n", subject, score)
	}

	// 零值 nil map：可读、可 len，写入会 panic
	var empty map[string]int
	fmt.Println("nil map 读:", empty["x"], "长度:", len(empty))
	// empty["x"] = 1 // panic: assignment to entry in nil map
}
