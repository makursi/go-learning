# AGENTS.md

Go 语言学习仓库：学习笔记、讲义代码、课后作业三类内容，每个示例与作业都能独立运行。

## 目录约定

| 目录 | 放什么 | 形态 |
| --- | --- | --- |
| `notes/` | 学习笔记，按主题编号分目录 | Markdown |
| `lectures/` | 课程讲义代码，每个示例一个目录 | `package main` + `main()` |
| `homework/` | 课后作业，每道题一个目录 | `package main` + `main()` |
| `projects/` | 综合练习项目（学到需要时再建） | 可含测试 |

- 路径全英文小写 slug；正文中文；术语首次出现写 `中文（English）`；代码注释中文。
- 单一根模块（`github.com/makursi/go-learning`），不给单个练习建 `go.mod`，暂不引入 `go.work`。
- 三边对齐靠同一个主题编号：`notes/02-data-structures/maps.md` ↔ `lectures/02-data-structures/maps/`；作业按题名走 `homework/02-data-structures/fruit-count/`，不要求和讲义同名。允许缺位，不强制每个知识点齐三份。
- 主题索引在 `notes/README.md`，新增知识点时同时更新它。

## 域文档

术语表是 `GLOSSARY.md`，决策记录在 `docs/adr/`，设计推导在 `docs/agents/design-notes.md`。

本仓**不建** `CONTEXT.md`，以此覆盖 `W/AGENTS.md` 对域文档入口的默认指向。

## Git

远端 `origin` = `github.com/makursi/go-learning`，默认分支 `main`。

`W/AGENTS.md` 把 `playgrounds/*` 列为「不是仓库，不做 git 操作」——那是过时描述，本仓有远端，按正常仓库处理。

**一次性 bootstrap 例外**：仓库原本是空仓、开不出 PR，所以首个提交直推 `main`。这**不构成先例**。此后所有改动：分支 → 提交 → `gh pr create` → 等确认 → `gh pr merge --squash --delete-branch`。

## 门禁

```sh
gofmt -l .        # 无输出即通过
go vet ./...
go build ./...
```

作业不配 `_test.go`，正确性靠 `go run` 看输出；只有 `projects/` 里的综合项目要求测试。

CI 暂缺。触发条件：第一次引入第三方依赖（Gin / GORM）时补 `.github/workflows/go.yml`，跑上面三条。
