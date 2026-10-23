package scaffold

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// status 是一次落点的结果，也是输出行里的那个词。
type status string

const (
	statusCreate status = "create"
	statusSkip   status = "skip"
)

// counts 给出这条结果在汇总行里的贡献。
func (s status) counts() (created, skipped int) {
	if s == statusCreate {
		return 1, 0
	}

	return 0, 1
}

// step 是规划层交给执行层的一条指令。
type step struct {
	rel     string // 仓库相对路径，用于输出
	abs     string // 绝对路径，用于读写
	content string
	exists  bool
}

// outcome 回答「此刻落盘会创建还是跳过」。
func (s step) outcome() status {
	if s.exists {
		return statusSkip
	}

	return statusCreate
}

// plan 是纯函数：除了问一句「这个路径存在吗」，它不碰文件系统。
// 规划与执行分开，是为了让 --dry-run 与真实运行共用同一份判断。
func plan(k kind, root, topic, name, title string, exists func(string) bool) step {
	rel := k.relPath(topic, name)
	abs := filepath.Join(root, filepath.FromSlash(rel))

	return step{
		rel:     rel,
		abs:     abs,
		content: k.content(title),
		exists:  exists(abs),
	}
}

// fileExists 是规划层的探测：只看路径在不在，不打开文件。
func fileExists(abs string) bool {
	_, err := os.Lstat(abs)

	return err == nil
}

// execute 把一条 step 落到磁盘上，返回实际发生的结果。
// 只创建不存在的文件：学习内容是手写的，工具不该有覆盖它的能力。
func execute(s step) (status, error) {
	if s.exists {
		return statusSkip, nil
	}

	if err := os.MkdirAll(filepath.Dir(s.abs), 0o755); err != nil {
		return statusCreate, fmt.Errorf("建目录 %s：%w", filepath.Dir(s.rel), err)
	}

	file, err := os.OpenFile(s.abs, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if errors.Is(err, fs.ErrExist) {
		return statusSkip, nil // 规划之后被别人抢先建了，按跳过算
	}
	if err != nil {
		return statusCreate, fmt.Errorf("建文件 %s：%w", s.rel, err)
	}

	if _, err := io.WriteString(file, s.content); err != nil {
		_ = file.Close()

		return statusCreate, fmt.Errorf("写文件 %s：%w", s.rel, err)
	}

	if err := file.Close(); err != nil {
		return statusCreate, fmt.Errorf("关文件 %s：%w", s.rel, err)
	}

	return statusCreate, nil
}
