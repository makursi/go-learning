package scaffold

import (
	"fmt"
	"regexp"
)

var (
	// 主题必须带两位编号：编号把笔记、讲义、作业三边对齐起来。
	topicPattern = regexp.MustCompile(`^[0-9]{2}-[a-z0-9]+(-[a-z0-9]+)*$`)
	// 名字是路径的一段，必须是英文小写 slug。
	slugPattern = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)
)

func validateTopic(topic string) error {
	if topicPattern.MatchString(topic) {
		return nil
	}

	return fmt.Errorf("主题 %q 不合规：应为两位编号加小写 slug，例如 03-functions", topic)
}

func validateName(name string) error {
	if slugPattern.MatchString(name) {
		return nil
	}

	return fmt.Errorf("名字 %q 不合规：应全小写，只用字母、数字和单个连字符，例如 fruit-count", name)
}
