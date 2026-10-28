// 题目：写 updateScore(scores map[string]int, subject string, delta int) 给某科加分、科目不存在则新建；再写 removeSubject 删除一科。
// 在 main 里验证改、增、删三种情况。

package main

import "fmt"

func updateScore(scores map[string]int, subject string, delta int) {
	scores[subject] += delta
}

func removeSubject(scores map[string]int, subject string) {
	delete(scores, subject)
}

func MapAddUpdateDelete() {
	scores := map[string]int{
		"语文": 90,
		"数学": 85,
		"英语": 92,
	}

	updateScore(scores, "英语", 30) // 改
	fmt.Println(scores)

	updateScore(scores, "物理", 10) // 增：科目不存在时等价于 0 + delta
	fmt.Println(scores)

	removeSubject(scores, "语文") // 删
	fmt.Println(scores)
}
