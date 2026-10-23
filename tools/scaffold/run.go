package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

// run 是唯一的接缝：参数、工作目录、两个输出流进，退出码出。
func run(args []string, cwd string, stdout, stderr io.Writer) int {
	if len(args) < 3 {
		fmt.Fprintf(stderr, "error: 用法：scaffold <kind> <topic> <name>\n")

		return exitUsage
	}

	root, err := findRoot(cwd)
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)

		return exitUsage
	}

	k, ok := lookupKind(args[0])
	if !ok {
		fmt.Fprintf(stderr, "error: 未知种类 %q\n", args[0])

		return exitUsage
	}

	if err := validateTopic(args[1]); err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)

		return exitUsage
	}

	if err := validateName(args[2]); err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)

		return exitUsage
	}

	rel := path.Join(k.root, args[1], args[2], k.entry)
	abs := filepath.Join(root, filepath.FromSlash(rel))

	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		fmt.Fprintf(stderr, "error: 建目录 %s：%v\n", filepath.Dir(rel), err)

		return exitIO
	}

	return createFile(abs, rel, mainGoTemplate, stdout, stderr)
}

// createFile 只创建不存在的文件：已存在就报告 skip 并保持原样。
// 学习内容是手写的，工具不该有覆盖它的能力。
func createFile(abs, rel, content string, stdout, stderr io.Writer) int {
	file, err := os.OpenFile(abs, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if errors.Is(err, fs.ErrExist) {
		fmt.Fprintf(stdout, "skip %s\n", rel)

		return exitOK
	}
	if err != nil {
		fmt.Fprintf(stderr, "error: 建文件 %s：%v\n", rel, err)

		return exitIO
	}

	if _, err := io.WriteString(file, content); err != nil {
		_ = file.Close()
		fmt.Fprintf(stderr, "error: 写文件 %s：%v\n", rel, err)

		return exitIO
	}

	if err := file.Close(); err != nil {
		fmt.Fprintf(stderr, "error: 关文件 %s：%v\n", rel, err)

		return exitIO
	}

	fmt.Fprintf(stdout, "create %s\n", rel)

	return exitOK
}
