package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// nextTopicNumber 扫笔记区的主题目录，给出「下一个可用编号」。
// 取现有最大编号 +1，而不是补空洞：编号是主题的身份，不该被回收。
func nextTopicNumber(root string) (string, bool) {
	entries, err := os.ReadDir(filepath.Join(root, "src", "notes"))
	if err != nil {
		if os.IsNotExist(err) {
			return "01", true
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

	return fmt.Sprintf("%02d", maxNumber+1), true
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
