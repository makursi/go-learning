//map 增删改：写 updateScore(scores map[string]int, subject string, delta int)，给某科加分；若科目不存在则新建。
// 再写 removeSubject(scores map[string]int, subject string) 删除一科。
// 在 main 里验证增、改、删后 map 的变化。

package main

func updateScore(scores map[string]int, subject string, delta int) {
	scores[subject] += delta
}

func removeSubject(scores map[string]int, subject string) {
	delete(scores, subject)
}
