package scaffold

// 退出码是工具的对外契约，脚本可能依赖它们，所以固定在这里。
const (
	exitOK    = 0 // 全部创建成功，或全部已存在而跳过
	exitIO    = 1 // 读写文件系统失败
	exitUsage = 2 // 用法或校验失败，什么都没写
)

// mainGoTemplate 是作业骨架：最小、能直接 go run 的 package main。
// 缩进必须是 tab、结尾必须有换行，否则生成的文件过不了 gofmt。
const mainGoTemplate = "package main\n\nfunc main() {\n}\n"

// noteContent 生成笔记骨架：只有一个一级标题和一个空行。
// 标题的缺省值由调用方先定好；这里不预置「要点」「TODO」这类段落，
// 写不出来的时候不该被逼着填。
func noteContent(title string) string {
	return "# " + title + "\n\n"
}

// goContent 是作业的内容生成器；--title 对它没有意义。
func goContent(string) string {
	return mainGoTemplate
}
