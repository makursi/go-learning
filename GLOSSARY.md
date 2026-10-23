# go-learning

一个 Go 语言学习仓库：记录语法理解、跟课敲的讲义代码、课后作业，以及逐步长出来的综合练习。

技术上只有一个根模块 `github.com/makursi/go-learning`，每个讲义示例和作业都放在自己的目录里，因此都能单独 `go run`，不会因为共用 `main()` 而冲突。

**内容分类、命名规则、Git 流程、门禁命令都写在 [AGENTS.md](AGENTS.md)。** 术语见 [GLOSSARY.md](GLOSSARY.md)，设计推导见 [docs/agents/design-notes.md](docs/agents/design-notes.md)。

## Language

### 构建与模块

**Module（模块）**：
由 `go.mod` 界定的版本化单元，是 Go 依赖管理的最小单位。本仓为 `github.com/makursi/go-learning`。
_Avoid_: 包、项目

**Package（包）**：
同一目录下 `package` 声明相同的一组源文件，是编译的最小单位。一个 Module 含多个 Package，反向不成立。
_Avoid_: 模块、库

**GOPATH**：
`go install` 的默认安装目标、以及模块缓存（`pkg\mod`）与校验库（`pkg\sumdb`）所在的根目录。本机为 `C:\Users\29634\go`。
_Avoid_: 工作区、项目目录

**GOBIN**：
`go install` 写出可执行文件的目录。本机为 `C:\Users\29634\go\bin`，已在 PATH 上。
_Avoid_: 安装目录、bin 目录

**GOPROXY**：
go 命令下载模块时查询的代理端点。本机固定为 `https://goproxy.cn,direct`（原因见 `docs/adr/0001-goproxy-mirror.md`）。
_Avoid_: 镜像源、源、仓库地址

### 工具链与编辑器

**gopls**：
Go 官方语言服务器，向编辑器提供补全、跳转、诊断与格式化。本机 v0.24.0。
_Avoid_: Go 插件、Go 语言服务

**LSP（Language Server Protocol）**：
编辑器与语言服务器之间的通用协议。gopls 是它在 Go 上的实现，VS Code 的 Go 扩展是它的客户端——两者缺一不可。
_Avoid_: 语法高亮插件

**Workspace（工作区）**：
本仓库根目录 `C:\Users\29634\Desktop\W\playgrounds\go-learning`。与 Go 的 `go.work` 工作区是两个无关概念。
_Avoid_: GOPATH

### 仓库分区

**Content（内容区）**：
`src/` 下被学习的全部内容——笔记、讲义、作业、项目。与 `tools/` 相对：内容区是「我在学的东西」，工具区是「服务于学习的程序」。`src/` 这个名字不指向 GOPATH，只是这个分区的容器。
_Avoid_: 源码、项目目录

**Tool（工具区）**：
`tools/` 下服务于学习仓库自身的程序，不是学习内容。目前含按主题生成目录与骨架的 `scaffold`。
_Avoid_: 脚本、辅助代码

**Scaffold（脚手架）**：
`tools/scaffold` 下的命令行工具，按「种类 + 主题编号 + 名字」在内容区建出目录与骨架文件。一次只落一边——笔记、讲义、作业三选一，不强制三边齐备，也不覆盖已有内容。
_Avoid_: 生成器、模板工具、init

### 内容分类

**Topic（主题）**：
笔记、讲义、作业共用的编号单元（如 `02-data-structures`）。三边靠它对齐，但允许缺位——不强制每个主题齐三份。
_Avoid_: 章节、单元

**Note（笔记）**：
`src/notes/` 下按 Topic 分目录的 Markdown，记录知识点本身。是复习入口，不承载完整可执行代码。
_Avoid_: 文档、教程

**Lecture（讲义）**：
`src/lectures/` 下跟课敲的示例代码，一个示例一个目录。演示某个概念的最小可运行片段，不是作业答案。
_Avoid_: 课件、demo

**Homework（作业）**：
`src/homework/` 下独立完成的题目，一题一目录。与 Lecture 的区别是「自己写的」而非「跟着敲的」，题目说明记在对应 Notes 里，不塞进作业目录。
_Avoid_: 练习、exercise

**Project（项目）**：
`src/projects/` 下的综合练习，有生命周期、可以长成独立仓，学到需要时再建。与 Lecture / Homework 的区别是它要求测试，而那两者不配 `_test.go`。
_Avoid_: 大作业、课设
