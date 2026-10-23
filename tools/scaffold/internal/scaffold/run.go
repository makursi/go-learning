package scaffold

import (
	"flag"
	"fmt"
	"io"
)

// options 是一次调用的可选项：种类无关的与种类特有的都放这里。
type options struct {
	title  string
	dryRun bool
}

// Run 是唯一的接缝：参数、工作目录、两个输出流进，退出码出。
// 测试与 main 都只经过它，其余都是它的内部。
func Run(args []string, cwd string, stdout, stderr io.Writer) int {
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
		reportTopicError(stderr, err, root)

		return exitUsage
	}

	if err := validateName(name); err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)

		return exitUsage
	}

	title := opts.title
	if k.titled && title == "" {
		title = name
	}

	item := plan(k, root, topic, name, title, fileExists)

	if opts.dryRun {
		report(stdout, item.outcome(), item.rel)

		return exitOK
	}

	result, err := execute(item)
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)

		return exitIO
	}

	report(stdout, result, item.rel)

	return exitOK
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

// reportTopicError 给主题类的用法错误补一句能直接照做的提示。
func reportTopicError(stderr io.Writer, cause error, root string) {
	if hint, ok := topicNumberHint(root); ok {
		fmt.Fprintf(stderr, "error: %v（%s）\n", cause, hint)

		return
	}

	fmt.Fprintf(stderr, "error: %v\n", cause)
}

// report 打印一条结果与一行汇总。
// dry-run 时两者都是计划值：只要真实运行不遇上 I/O 失败，输出逐字相同。
func report(stdout io.Writer, s status, rel string) {
	created, skipped := s.counts()

	fmt.Fprintf(stdout, "%s %s\n", s, rel)
	fmt.Fprintf(stdout, "%d created, %d skipped\n", created, skipped)
}
