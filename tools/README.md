# tools/

工具区服务于学习仓库自身，与 `src/` 的内容隔离。用 `go run` 直接跑，不要安装。

## scaffold：按主题生成目录与骨架

生成一条笔记：

```sh
go run ./tools/scaffold note 03-functions closures
```

指定标题（缺省用名字）：

```sh
go run ./tools/scaffold note --title "闭包（Closures）" 03-functions closures
```

生成讲义与作业：

```sh
go run ./tools/scaffold lecture  03-functions closures
go run ./tools/scaffold homework 03-functions fruit-count
```

先预览、再落盘：

```sh
go run ./tools/scaffold note --dry-run 03-functions closures
```

记住这五条：

1. 把 flag 写在位置参数之前。
2. 给主题写全两位编号（`03-functions`）。写漏或写错，报错会告诉你下一个可用编号。
3. 名字用小写 slug：字母、数字、单个连字符。
4. 三边按需各自生成，不要为了对齐硬凑。
5. 不要指望覆盖——没有 `--force`，已存在的文件一律 `skip`。

退出码：`0` 成功或跳过；`1` 读写文件系统失败；`2` 用法或校验失败，且什么都没写。

完整说明见 [scaffold/README.md](scaffold/README.md)。

## 加新工具

保持零第三方依赖：根目录放一个薄 `main`，核心放 `internal/`，测试放 `test/`。动手前先跟用户确认。
