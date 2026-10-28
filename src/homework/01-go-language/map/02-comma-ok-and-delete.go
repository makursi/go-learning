// 题目：给定 scores := map[string]int{"语文": 90, "数学": 0}，分别判断「语文」「数学」「英语」是否存在，并正确打印成绩或「不存在」；
// 再删除「数学」，用 len 打印剩余键值对数量。

package main

import "fmt"

func CommaOkAndDelete() {
	scores := map[string]int{
		"语文": 90,
		"数学": 0,
	}

	if value, ok := scores["语文"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}

	if value, ok := scores["数学"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}

	if value, ok := scores["英语"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}

	delete(scores, "数学")

	fmt.Println(len(scores))
}
