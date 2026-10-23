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

// options 是一次调用的可选项：与种类无关的放在这里，种类特有的也放这里。
type options struct {
	title  string
	dryRun bool
}

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

	opts, rest, usage := parseKindArgs(k, args[1:], stderr)
	if usage {
		return exitUsage
	}

	if len(rest) != 2 {
		fmt.Fprintf(stderr, "error: %s 需要 <topic> <name> 两个参数，拿到 %d 个；用法：%s\n", k.name, len(rest), k.usageLine())

		return exitUsage
	}

	topic, name := rest[0], rest[1]

	root, err := findRoot(cwd)
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)

		return exitUsage
	}

	if err := validateTopic(topic); err != nil {
		if next, ok := nextTopicNumber(root); ok {
			fmt.Fprintf(stderr, "error: %v（下一个可用编号是 %s）\n", err, next)
		} else {
			fmt.Fprintf(stderr, "error: %v\n", err)
		}

		return exitUsage
	}

	if err := validateName(name); err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)

		return exitUsage
	}

	rel := k.relPath(topic, name)
	abs := filepath.Join(root, filepath.FromSlash(rel))

	if opts.dryRun {
		report(stdout, planStatus(abs), rel)

		return exitOK
	}

	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		fmt.Fprintf(stderr, "error: 建目录 %s：%v\n", filepath.Dir(rel), err)

		return exitIO
	}

	status, code := createFile(abs, rel, k.content(name, opts.title), stdout, stderr)
	if code != exitOK {
		return code
	}

	report(stdout, status, rel)

	return exitOK
}

const (
	statusCreate = "create"
	statusSkip   = "skip"
)

// report 打印一条结果与一行汇总。
// dry-run 时这两个值都是计划值，与真实运行逐字一致，方便直接对比。
func report(stdout io.Writer, status, rel string) {
	created, skipped := 0, 1
	if status == statusCreate {
		created, skipped = 1, 0
	}

	fmt.Fprintf(stdout, "%s %s\n", status, rel)
	fmt.Fprintf(stdout, "%d created, %d skipped\n", created, skipped)
}

// parseKindArgs 按种类解析它自己的 flag 集合：不接受 --title 的种类传了它，
// 由标准库报「flag provided but not defined」，不用手写分支。
func parseKindArgs(k kind, args []string, stderr io.Writer) (options, []string, bool) {
	flags := flag.NewFlagSet(k.name, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var opts options
	flags.BoolVar(&opts.dryRun, "dry-run", false, "只打印计划，不落盘")

	if k.titled {
		flags.StringVar(&opts.title, "title", "", "一级标题，缺省用名字")
	}

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(stderr, "error: %v\n用法：%s\n", err, k.usageLine())

		return options{}, nil, true
	}

	return opts, flags.Args(), false
}

// planStatus 回答「如果不带 --dry-run，这一步会创建还是跳过」。
func planStatus(abs string) string {
	if _, err := os.Lstat(abs); err == nil {
		return statusSkip
	}

	return statusCreate
}

// createFile 只创建不存在的文件：已存在就报告 skip 并保持原样。
// 学习内容是手写的，工具不该有覆盖它的能力。
func createFile(abs, rel, content string, stdout, stderr io.Writer) (string, int) {
	file, err := os.OpenFile(abs, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if errors.Is(err, fs.ErrExist) {
		return statusSkip, exitOK
	}
	if err != nil {
		fmt.Fprintf(stderr, "error: 建文件 %s：%v\n", rel, err)

		return "", exitIO
	}

	if _, err := io.WriteString(file, content); err != nil {
		_ = file.Close()
		fmt.Fprintf(stderr, "error: 写文件 %s：%v\n", rel, err)

		return "", exitIO
	}

	if err := file.Close(); err != nil {
		fmt.Fprintf(stderr, "error: 关文件 %s：%v\n", rel, err)

		return "", exitIO
	}

	return statusCreate, exitOK
}
