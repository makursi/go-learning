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

// topicNumber 取主题名开头的两位编号；不是「两位数字 + 连字符」就返回 false。
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
