package main

// 退出码是工具的对外契约，脚本可能依赖它们，所以固定在这里。
const (
	exitOK    = 0 // 全部创建成功，或全部已存在而跳过
	exitIO    = 1 // 读写文件系统失败
	exitUsage = 2 // 用法或校验失败，什么都没写
)

// mainGoTemplate 是讲义与作业共用的最小可运行骨架。
// 缩进必须是 tab、结尾必须有换行，否则生成的文件过不了 gofmt。
const mainGoTemplate = "package main\n\nfunc main() {\n}\n"
