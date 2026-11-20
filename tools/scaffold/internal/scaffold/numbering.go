package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// topicNumberHint 给主题编号的用法错误补一句可执行的提示。
// 编号只有两位：用满 99 之后不再给数字，否则提示值自己就过不了校验。
func topicNumberHint(root string) (string, bool) {
	entries, err := os.ReadDir(filepath.Join(root, "src", "notes"))
	if err != nil {
		if os.IsNotExist(err) {
			return "下一个可用编号是 01", true
		}

		return "", false
	}

	maxNumber := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if number, ok := topicNumber(entry.Name()); ok && number > maxNumber {
			maxNumber = number
		}
	}

	if maxNumber >= 99 {
		return "两位编号已用尽，新主题需要先放宽编号规则", true
	}

	return fmt.Sprintf("下一个可用编号是 %02d", maxNumber+1), true
}

// nextExerciseNumber 给出主题目录下一个可用的两位题号。
// 题号是目录内的追加顺序，所以取「现有最大 + 1」；目录读不出来时按「还没有题」
// 处理——真正的失败会在建目录那一步带着路径报出来，这里不该先把它报成用法错误。
func nextExerciseNumber(root, topic string) (string, error) {
	entries, err := os.ReadDir(filepath.Join(root, "src", "homework", topic))
	if err != nil {
		return "01", nil
	}

	maxNumber := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if number, ok := topicNumber(entry.Name()); ok && number > maxNumber {
			maxNumber = number
		}
	}

	if maxNumber >= 99 {
		return "", fmt.Errorf("题号已用尽：%s 下已有 99，两位编号放不下更多题", topic)
	}

	return fmt.Sprintf("%02d", maxNumber+1), nil
}

// topicNumber 取名字开头的两位编号；不是「两位数字 + 连字符」就返回 false。
func topicNumber(name string) (int, bool) {
	if len(name) < 4 || name[2] != '-' {
		return 0, false
	}

	number, err := strconv.Atoi(name[:2])
	if err != nil || number <= 0 {
		return 0, false
	}

	return number, true
}
