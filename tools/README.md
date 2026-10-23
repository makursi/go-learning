# tools/

服务于学习仓库自身的程序，与 `src/` 的内容隔离。每个工具是一个独立 `package main`，零第三方依赖。

| 工具 | 做什么 |
| --- | --- |
| [`scaffold`](scaffold/README.md) | 按主题编号生成笔记、讲义、作业的目录与骨架 |

日常用 `go run ./tools/<name> ...` 跑，不必安装。往这里加新工具之前先跟用户确认。
