package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// run 是唯一的接缝：参数、工作目录、两个输出流进，退出码出。
func run(args []string, cwd string, stdout, stderr io.Writer) int {
	if len(args) < 1 {
		fmt.Fprintf(stderr, "error: 用法：scaffold <kind> <topic> <name>\n")

		return exitUsage
	}

	k, ok := lookupKind(args[0])
	if !ok {
		fmt.Fprintf(stderr, "error: 未知种类 %q\n", args[0])

		return exitUsage
	}

	title, rest, usage := parseKindArgs(k, args[1:], stderr)
	if usage {
		return exitUsage
	}

	if len(rest) != 2 {
		fmt.Fprintf(stderr, "error: %s 需要 <topic> <name> 两个参数，拿到 %d 个；用法：%s\n", k.name, len(rest), k.usageLine())

		return exitUsage
	}

	topic, name := rest[0], rest[1]

	if err := validateTopic(topic); err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)

		return exitUsage
	}

	if err := validateName(name); err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)

		return exitUsage
	}

	root, err := findRoot(cwd)
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)

		return exitUsage
	}

	rel := k.relPath(topic, name)
	abs := filepath.Join(root, filepath.FromSlash(rel))

	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		fmt.Fprintf(stderr, "error: 建目录 %s：%v\n", filepath.Dir(rel), err)

		return exitIO
	}

	return createFile(abs, rel, k.content(name, title), stdout, stderr)
}

// parseKindArgs 按种类解析它自己的 flag 集合：不接受 --title 的种类传了它，
// 由标准库报「flag provided but not defined」，不用手写分支。
func parseKindArgs(k kind, args []string, stderr io.Writer) (title string, rest []string, usage bool) {
	flags := flag.NewFlagSet(k.name, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	if k.titled {
		flags.StringVar(&title, "title", "", "一级标题，缺省用名字")
	}

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(stderr, "error: %v\n用法：%s\n", err, k.usageLine())

		return "", nil, true
	}

	return title, flags.Args(), false
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
