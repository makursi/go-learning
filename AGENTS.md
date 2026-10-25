# AGENTS.md

Go 语言学习仓库：学习笔记、讲义代码、课后作业三类内容，每个示例与作业都能独立运行。

## 目录约定

| 目录 | 放什么 | 形态 |
| --- | --- | --- |
| `src/` | **内容区**：被学习的全部内容，下含三类 | — |
| `src/notes/` | 学习笔记，按主题编号分目录 | Markdown |
| `src/lectures/` | 课程讲义代码，每个示例一个目录 | `package main` + `main()` |
| `src/homework/` | 课后作业，每道题一个目录 | `package main` + `main()` |
| `src/projects/` | 综合练习项目（学到需要时再建） | 可含测试 |
| `tools/` | **工具区**：服务于学习本身的程序 | 独立 `package main` |
| `docs/` | 设计推导与决策记录 | Markdown |
| `docs/specs/` | 已发布 spec 的仓库内快照，以 issue 为准 | Markdown |

**`src/` 是刻意保留的，不要拍平。** 它把「我在学的东西」（`src/`）和「服务于学习的工具、文档、仓库配置」（`tools/` `docs/` 及根文件）分成两个区。这一点不符合 Go 的一般习惯——`src/` 在 GOPATH 时代专指 `$GOPATH/src/<import-path>`——所以当初讨论时被建议去掉。保留是用户的决定，理由是可维护性：内容与工具混在根级会分不清谁属于谁。取舍全过程记在 `docs/agents/design-notes.md`。

已知代价：将来从这些内容里长出的可复用包，导入路径是 `github.com/makursi/go-learning/src/...`，比无 `src/` 时多一级。

`tools/` 放服务于学习本身的程序，目前只有一个：`scaffold`，按主题生成目录与骨架（见 `tools/scaffold/README.md`）。**往 `tools/` 里加东西之前先跟用户确认。**

- 路径全英文小写 slug；正文中文；术语首次出现写 `中文（English）`；代码注释中文。
- 单一根模块（`github.com/makursi/go-learning`），不给单个练习建 `go.mod`，暂不引入 `go.work`。
- 三边对齐靠同一个主题编号：`src/notes/02-data-structures/maps.md` ↔ `src/lectures/02-data-structures/maps/`；作业按题名走 `src/homework/02-data-structures/fruit-count/`，不要求和讲义同名。允许缺位，不强制每个知识点齐三份。
- 目录与骨架由 `go run ./tools/scaffold ...` 生成，不手搓：`note` / `lecture` / `homework` 三条命令，flag 写在位置参数之前。

## 域文档

术语表是 `GLOSSARY.md`，决策记录在 `docs/adr/`，设计推导在 `docs/agents/design-notes.md`。

本仓**不建** `CONTEXT.md`，以此覆盖 `W/AGENTS.md` 对域文档入口的默认指向。

## Agent skills

### Issue tracker

Issue 与 spec 发布在 `makursi/go-learning` 的 GitHub Issues，一律用 `gh` CLI 操作；仓库内 `docs/specs/` 留一份快照，冲突时以 issue 为准。见 `docs/agents/issue-tracker.md`。

### Triage labels

五个 triage 角色与同名标签 1:1 对应；`ready-for-agent` 表示 spec 已定、可直接实现。见 `docs/agents/triage-labels.md`。

### Domain docs

单上下文：根目录 `GLOSSARY.md` + `docs/adr/`（另见 `docs/agents/design-notes.md`）。见 `docs/agents/domain.md`。

## Git

远端 `origin` = `github.com/makursi/go-learning`，默认分支 `main`。

`W/AGENTS.md` 把 `playgrounds/*` 列为「不是仓库，不做 git 操作」——那是过时描述，本仓有远端，按正常仓库处理。

**一次性 bootstrap 例外**：仓库原本是空仓、开不出 PR，所以首个提交直推 `main`。这**不构成先例**。

此后所有改动：分支 → 提交 → `gh pr create` → **等用户确认** → `gh pr merge --squash --delete-branch`。这条对本仓是用户明确要求保留的，别以"单人仓没必要"为由省掉。

手动走这套流程时，照 `docs/pr-workflow.md` 的七步做。

## 门禁

```sh
gofmt -l .        # 无输出即通过
go vet ./...
go build ./...
```

作业不配 `_test.go`，正确性靠 `go run` 看输出；`src/projects/` 里的综合项目与 `tools/` 里的程序要配测试。

CI 暂缺。触发条件：第一次引入第三方依赖（Gin / GORM）时补 `.github/workflows/go.yml`，跑上面三条。
