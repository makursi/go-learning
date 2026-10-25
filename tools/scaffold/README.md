# scaffold

按学习内容的结构生成目录与骨架。一次调用只落一个种类的一个条目：笔记与作业互不牵连。

## 命令

```sh
go run ./tools/scaffold note     --title "映射（Maps）" 02-data-structures maps
go run ./tools/scaffold homework 02-data-structures fruit-count
```

flag 必须写在位置参数之前——标准库 `flag` 在第一个位置参数处就停止解析，`note 02-data-structures maps --title x` 会把 `--title` 丢掉。

| 种类 | 落点 | `--title` |
| --- | --- | --- |
| `note` | `src/notes/<topic>/<name>.md` | 有效，缺省用名字 |
| `homework` | `src/homework/<topic>/<name>/main.go` | 不支持 |

- `<topic>` 必须带两位编号，形如 `03-functions`。编号写错或不写时，报错里会给出下一个可用编号。
- `<name>` 是英文小写 slug：字母、数字、单个连字符。
- 只建该种类用到的那棵树，不会顺手在另外两棵树里造空目录。
- 笔记只有一个一级标题和一个空行；作业是最小的 `package main` + `main()`。作业目录不放 `README.md`，题面记在对应笔记里。

## 代码结构

```text
tools/scaffold/
├── main.go              # 入口：取 cwd，调 Run，透传退出码
├── internal/scaffold/   # 核心：种类表、规划、执行、校验、编号、模板
└── test/                # 外部测试包，只经 Run 进入
```

## 选项

- `--dry-run` 只打印计划，不落盘。输出的 `create` / `skip` 行与计数都是计划值，与真实运行逐字一致。
- 没有 `--force`：已存在的文件一律 `skip`，要覆盖就手动删。

## 输出与退出码

```text
$ go run ./tools/scaffold note 02-data-structures maps
create src/notes/02-data-structures/maps.md
1 created, 0 skipped
```

| 退出码 | 含义 |
| --- | --- |
| 0 | 全部创建成功，或目标已存在而跳过 |
| 1 | 读写文件系统失败 |
| 2 | 用法或校验失败（参数个数、slug、编号、不在本仓库内），什么都没写 |

错误一律走 stderr，带 `error:` 前缀。

## 安装（可选）

```sh
go install ./tools/scaffold
```

装进 `GOBIN`（本机已在 PATH 上），之后任何目录都能直接敲 `scaffold`。

## 加一种内容种类

`internal/scaffold/kinds.go` 的表里加一行 + `template.go` 里加一个内容生成器，判定逻辑不用动。将来真要加一类，先把理由写进 `docs/agents/design-notes.md`——这个仓刻意只有两类。
