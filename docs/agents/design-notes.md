# 设计笔记：go-learning

## 问题

把 `C:\Users\29634\Desktop\W\playgrounds\go-learning` 建成一个长期可维护的 Go 学习仓库：三类内容（学习笔记 / 课程讲义代码 / 课后作业）各自独立、可单独运行、可互相检索，并与远端 `makursi/go-learning` 保持同步。

起点的实际状态（2026-10-23 核实）：本地 git 已 init、零提交、分支 `master`、无 remote；远端 public 且为空仓、默认分支 `main`；`go.mod` 为 `module go-start`；`README.md` 是 552 行的仓库设计建议书而非仓库说明；根目录有一个空的 `package main`（`index.go`）；`notes/` `lectures/` `projects/` 为空目录；`GLOSSARY.md` 是上一轮实验残留，标题与路径都写作 `go-start`。

## 已定分支

### 仓库位置

留在 `playgrounds/go-learning`，不迁往 `apps/`。它是学习档案而非应用，与 `playgrounds/rn-start` 同类同位置。

已知风险：`W/AGENTS.md` 把 `playgrounds/*` 列为「不是仓库，不做 git 操作」。缓解手段是仓内 `AGENTS.md` —— 工作区文件自己声明「各项目的约定由该项目自己制定，动手前先读当前仓的 AGENTS.md」。

### 模块身份

`module go-start` → `github.com/makursi/go-learning`。仓库已是 public，将来从笔记里长出的可复用包需要与远端一致的导入路径。

### 单一根模块

一个根 `go.mod` 管全部内容，不给单个练习建 `go.mod`，暂不引入 `go.work`。

收益：根目录一条 `go build ./...` 就能验证全部代码，依赖集中升级。代价：各练习共享同一份依赖声明，两个练习要不同版本同一库时会打架 —— 学习阶段几乎不可能发生，真发生再拆，拆比提前铺 `go.work` 便宜。

### 顶层分类

`notes/` / `lectures/` / `homework/` 三类为常设；`projects/` 为第四类但**按需创建** —— 它是另一种东西（有生命周期、可能长成独立仓），提前建空目录只是制造假进度。

### 三边对齐（宽松）

笔记、讲义、作业靠同一个主题编号对齐，**允许缺位**，不强制每个知识点齐三份。

命名规则：`notes/02-data-structures/maps.md` ↔ `lectures/02-data-structures/maps/` 同名；作业按题名走 `homework/02-data-structures/fruit-count/`，不要求和讲义同名。索引在 `notes/README.md`，每个知识点一行链到三边，缺的留空 —— 这是宽松对齐能成立的前提。

强制齐三份会在「这个知识点我写不出作业」时逼着造垃圾。

### 命名与语言

路径全英文小写 slug，正文中文，术语首次出现写 `中文（English）`，代码注释中文。

理由：路径要过终端、URL、Go 工具链和编辑器搜索，中文路径在 Windows 上的坑不值得为省一次翻译去踩；正文是给自己复习的，中文读得快。

### 域文档入口

保留 `GLOSSARY.md` 作为唯一术语表，不新增 `CONTEXT.md`，并在仓内 `AGENTS.md` 声明「域文档 = `GLOSSARY.md` + `docs/adr/`」，覆盖 `W/AGENTS.md` 对 `CONTEXT.md` 的默认指向。

理由：术语表与「这个仓是什么」是两类内容，合并后两边都会脏；且工作区文件明确授权项目自定约定。

### 作业验收

作业不配 `_test.go`，也不写批量运行脚本。用 `gofmt -l .` / `go vet ./...` / `go build ./...` 保证「能编译、没毛病」，正确性靠手动 `go run` 看输出。题面记在对应 Notes 里，不在作业目录塞 `README.md`。

学习期的作业答案是一次性的，写测试等于给草稿纸配档案柜；只有 `projects/` 里的综合项目要求测试。

### 编辑器

建 `.vscode/extensions.json` 只推荐 `golang.go`；不建 `settings.json` —— `gopls v0.24.0` 已在 `GOBIN` 且上了 PATH。

注意 `.gitignore` 里 `.vscode/*` 与 `!.vscode/extensions.json` 必须成对出现，否则推荐文件自己也被忽略。

### 首次同步（bootstrap 例外）

`git branch -m master main` → 接上 `origin` → 首个提交直推 `main`。

理由：远端是空仓，没有 base 分支就开不出 PR。这只是一次性 bootstrap，**不构成先例**；此后所有改动一律走 分支 → PR → 等确认 → squash 合并。已写进仓内 `AGENTS.md`，避免下次被误引。

### 根目录内容

根目录最终只有 `go.mod` / `README.md` / `.gitignore` / `AGENTS.md` / `GLOSSARY.md` / `.vscode/` / `docs/` / 三类内容目录。

- `index.go`（空的 `package main`）删除：根目录不该有 `main` 包。
- 空目录不补 `.gitkeep`：git 本就不跟踪空目录，逼着目录只在有内容时出现。
- `README.md` 的 552 行原文**不归档**，留作本次工作的参考，收尾时替换为单行标题，正文由用户自己编写。

### CI

暂不加 GitHub Actions，留明确触发条件：第一次引入第三方依赖（Gin / GORM）时补 `.github/workflows/go.yml`，跑 `gofmt -l` / `go vet` / `go build`。在那之前，单人学习仓的 CI 只是在为一个不存在的 PR 流程站岗。

## 待定分支

无。前沿已空。

## 已否分支

- **迁到 `apps/go-learning`**：它是学习档案不是应用；`playgrounds/rn-start` 已确立「playgrounds 下可以有真仓」的先例，迁移只会换来一次无收益的路径改写。
- **模块路径保持 `go-start`**：无法被 `go get`；仓库已 public，改用远端路径的成本只有一行。
- **改用 `CONTEXT.md` 迎合工作区默认**：丢掉了「术语表只放术语」这条边界。
- **每个练习独立 `go.mod` / 提前上 `go.work`**：多出成倍的 `go.mod`+`go.sum` 要维护，换来的是学习阶段用不到的多版本隔离。
- **强制每个知识点齐三份（笔记 / 讲义 / 作业）**：写不出作业的知识点会被迫造假内容。
- **顶层现在就建 `projects/` 空目录**：空目录进不了 git，等于只建了个心理安慰。
- **作业配 `_test.go` 或批量运行脚本**：一次性答案不值得配测试；批量运行脚本在 `go run ./...` 本身不支持的前提下要自己解析目录，维护成本高于收益。
- **现在加 CI**：见「CI」小节，改为按触发条件延后。
- **为单根模块 / bootstrap 例外 / 模块改名写 ADR**：分别不满足「难以逆转」「不解释会困惑」「真有权衡」——单根模块可逆，bootstrap 是事实约束不是权衡，改名是纠正错字。
- **把首个提交推成分支再手动合 PR**：空仓没有 base 分支，PR 开不出来，属于把流程仪式化。
- **`README.md` 原文移入 `docs/reference/` 存档**：那 552 行是一次性设计建议，结论已进本文件，留着只是第二份会过期的文档。
