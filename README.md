# go-learning

个人Golang学习记录

## 文档

### 官方 · 参考

- [Go 语言规范](https://go.dev/ref/spec) — 语法、类型、求值的最终解释权
- [标准库索引](https://pkg.go.dev/std) — 全部标准库包的 API
- [Go 命令与工具](https://go.dev/doc/cmd) — `go` 命令及自带工具的完整用法
- [Go Modules 参考](https://go.dev/ref/mod) — 模块、版本、代理、校验的完整规则
- [go.mod 参考](https://go.dev/doc/modules/gomod-ref) — `go.mod` 每条指令的含义
- [内存模型](https://go.dev/ref/mem) — 什么情况下一个 goroutine 能看到另一个的写
- [版本发布历史](https://go.dev/doc/devel/release) — 每个版本改了什么
- [Go 博客](https://go.dev/blog) — 切片内部结构、反射三定律、defer 的实现都在这里
- [FAQ](https://go.dev/doc/faq) — 常见问题的官方回答
- [Go Wiki](https://go.dev/wiki) — 社区维护的官方 wiki

### 官方 · 学习

- [A Tour of Go](https://go.dev/tour/) — 四段式交互教程，含并发与练习
- [Effective Go](https://go.dev/doc/effective_go) — 怎么写出地道的 Go
- [如何写 Go 代码](https://go.dev/doc/code) — 包、模块、测试的最小工作流
- [教程：创建模块](https://go.dev/doc/tutorial/create-module) — 函数、错误、数组、map、单测一条龙
- [教程：多模块工作区](https://go.dev/doc/tutorial/workspaces) — `go.work` 到底解决什么问题
- [教程：泛型入门](https://go.dev/doc/tutorial/generics)
- [教程：模糊测试](https://go.dev/doc/tutorial/fuzz)
- [教程：用 Gin 写 REST API](https://go.dev/doc/tutorial/web-service-gin)
- [教程：访问关系型数据库](https://go.dev/doc/tutorial/database-access)
- [文档注释规范](https://go.dev/doc/comment) — `//` 注释如何生成 `go doc`
- [垃圾回收指南](https://go.dev/doc/gc-guide) — GC 怎么工作、怎么调
- [诊断工具](https://go.dev/doc/diagnostics) — 出问题该用哪个工具查
- [数据竞争检测](https://go.dev/doc/articles/race_detector) — `-race` 的手册
- [性能剖析](https://go.dev/blog/pprof) — pprof 的使用与读图
- [覆盖率](https://go.dev/doc/build-cover) — `-cover` 与覆盖率报告
- [PGO 入门](https://go.dev/doc/pgo) — 按运行剖面优化编译
- [Go by Example](https://gobyexample.com/) — 英文示例集，每个知识点一段可运行代码（第三方）

### 中文

- [Go 编程圣典](https://go-docs.pages.dev/) — 个人系统化中文教程，按主题成篇（非官方）
- [golang.google.cn](https://golang.google.cn/) — `go.dev` 的国内镜像，内容同源，仍是英文

### 生态 · 框架与库

- [Gin](https://gin-gonic.com/docs/) — 最常用的 Web 框架
- [Echo](https://echo.labstack.com/docs) — 路由与中间件风格更接近标准库
- [Fiber](https://docs.gofiber.io/) — 基于 fasthttp，吞吐优先
- [chi](https://github.com/go-chi/chi) — 纯 `net/http` 风格的路由器
- [GORM](https://gorm.io/docs/) — 最常用的 ORM
- [sqlc](https://docs.sqlc.dev/) — 由 SQL 生成类型安全的 Go 代码，不走反射
- [cobra](https://cobra.dev/) — 命令行应用骨架
- [viper](https://github.com/spf13/viper) — 配置读取（文件、环境变量、远程）
- [zap](https://github.com/uber-go/zap) — 结构化日志
- [testify](https://github.com/stretchr/testify) — 断言与 mock
- [wire](https://github.com/google/wire) — 编译期依赖注入

### 生态 · 工具链

- [gopls](https://go.dev/gopls) — 官方语言服务器，编辑器的补全与跳转全靠它
- [Delve](https://github.com/go-delve/delve) — 调试器，`dlv debug` 的底层
- [staticcheck](https://staticcheck.dev/) — 比 `go vet` 更狠的静态检查
- [golangci-lint](https://golangci-lint.run/) — 把多种 linter 合成一条命令
- [air](https://github.com/air-verse/air) — 改动自动重编译重启
- [Go Playground](https://go.dev/play) — 在线运行片段，贴代码问问题用它

### 生态 · 社区与资讯

- [Go 源码](https://github.com/golang/go) — 读标准库实现的地方
- [Go Issue Tracker](https://github.com/golang/go/issues) — 语言与工具链的问题都在这里
- [awesome-go](https://github.com/avelino/awesome-go) — 按领域分类的库清单
- [Golang Weekly](https://golangweekly.com/) — 每周生态动态
- [r/golang](https://www.reddit.com/r/golang/)
- [Gopher Slack](https://invite.slack.golangbridge.org/)
