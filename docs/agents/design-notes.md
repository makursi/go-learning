# 设计笔记：go-learning

## 问题

把 `C:\Users\29634\Desktop\W\playgrounds\go-learning` 建成一个长期可维护的 Go 学习仓库：三类内容（学习笔记 / 课程讲义代码 / 课后作业）各自独立、可单独运行、可互相检索，并与远端 `makursi/go-learning` 保持同步。

起点的实际状态（2026-10-23 核实）：本地 git 已 init、零提交、分支 `master`、无 remote；远端 public 且为空仓、默认分支 `main`；`go.mod` 为 `module go-start`；`README.md` 是 552 行的仓库设计建议书而非仓库说明；根目录有一个空的 `package main`（`index.go`）；`notes/` `lectures/` `projects/` 为空目录；`GLOSSARY.md` 是上一轮实验残留，标题与路径都写作 `go-start`。

## 已定分支

### 仓库位置

留在 `playgrounds/go-learning`，不迁往 `apps/`。它是学习档案而非应用，与 `playgrounds/rn-start` 同类同位置。

已知风险：`W/AGENTS.md` 把 `playgrounds/*` 列为「不是仓库，不做 git 操作」。缓解手段是仓内 `AGENTS.md` —— 工作区文件自己声明「各项目的约定由该项目自己制定，动手前先读当前仓的 AGENTS.md」。

### 内容区 / 工具区分离（`src/` 与 `tools/`）

全部学习内容放在 `src/` 之下：`src/notes/` `src/lectures/` `src/homework/`，将来还有 `src/projects/`。服务于学习本身的程序放在 `tools/`（目前只有占位文件，预留给一个自动创建目录结构的 CLI）。

用户给的理由：没有 `src/` 这一层，学习内容会和 `tools/` 平铺在根目录，分不清谁属于谁，可维护性变差。`src/` 这个名字**不指向 GOPATH 的 `$GOPATH/src`**，它只是「内容区」这个分区的容器名。

这项决定覆盖了此前的「三类放根级」方案（见下方已否分支）。已知代价：将来从内容里长出的可复用包，导入路径是 `github.com/makursi/go-learning/src/...`，比无 `src/` 时多一级。

因为不符合 Go 的一般习惯，`AGENTS.md` 里专门写了一条「`src/` 是刻意保留的，不要拍平」，防止后来的 agent 看到就顺手搬走。

### 模块身份

`module go-start` → `github.com/makursi/go-learning`。仓库已是 public，将来从笔记里长出的可复用包需要与远端一致的导入路径。

### 单一根模块

一个根 `go.mod` 管全部内容（含 `src/` 与 `tools/`），不给单个练习建 `go.mod`，暂不引入 `go.work`。

收益：根目录一条 `go build ./...` 就能验证全部代码，依赖集中升级。代价：各练习共享同一份依赖声明，两个练习要不同版本同一库时会打架 —— 学习阶段几乎不可能发生，真发生再拆，拆比提前铺 `go.work` 便宜。

### 三边对齐（宽松）

笔记、讲义、作业靠同一个主题编号对齐，**允许缺位**，不强制每个知识点齐三份。

命名规则：`src/notes/02-data-structures/maps.md` ↔ `src/lectures/02-data-structures/maps/` 同名；作业按题名走 `src/homework/02-data-structures/fruit-count/`，不要求和讲义同名。索引在 `src/notes/README.md`，每个知识点一行链到三边，缺的留空 —— 这是宽松对齐能成立的前提。

强制齐三份会在「这个知识点我写不出作业」时逼着造垃圾。

### 命名与语言

路径全英文小写 slug，正文中文，术语首次出现写 `中文（English）`，代码注释中文。

理由：路径要过终端、URL、Go 工具链和编辑器搜索，中文路径在 Windows 上的坑不值得为省一次翻译去踩；正文是给自己复习的，中文读得快。

### 域文档入口

保留 `GLOSSARY.md` 作为唯一术语表，不新增 `CONTEXT.md`，并在仓内 `AGENTS.md` 声明「域文档 = `GLOSSARY.md` + `docs/adr/`」，覆盖 `W/AGENTS.md` 对 `CONTEXT.md` 的默认指向。

理由：术语表与「这个仓是什么」是两类内容，合并后两边都会脏；且工作区文件明确授权项目自定约定。

### 作业验收

作业不配 `_test.go`，也不写批量运行脚本。用 `gofmt -l .` / `go vet ./...` / `go build ./...` 保证「能编译、没毛病」，正确性靠手动 `go run` 看输出。题面记在对应 Notes 里，不在作业目录塞 `README.md`。

学习期的作业答案是一次性的，写测试等于给草稿纸配档案柜；只有 `src/projects/` 里的综合项目要求测试。

**空模块例外**：还没有任何 Go 包时 `go vet ./...` 会以 `no packages to vet` 退出 1。这不是失败，已写进 `AGENTS.md`。

### 编辑器

建 `.vscode/extensions.json` 只推荐 `golang.go`；不建 `settings.json` —— `gopls v0.24.0` 已在 `GOBIN` 且上了 PATH。

注意 `.gitignore` 里 `.vscode/*` 与 `!.vscode/extensions.json` 必须成对出现，否则推荐文件自己也被忽略。

### 空目录与占位

顶层 `src/notes/` `src/lectures/` `src/homework/` 保留为空目录（git 不跟踪，只在本地存在）。`tools/` 因为空目录 clone 后会消失，放了一个 `tools/README.md` 占位，说明它预留给什么。

### Git 流程（保留 PR）

首次同步是一次性 bootstrap 例外：`git branch -m master main` → 接上 `origin` → 首个提交直推 `main`。这**不构成先例**。

此后所有改动一律 分支 → 提交 → PR → **等用户确认** → squash 合并并删分支。用户明确选择保留这条流程，理由是「分支 → PR → 自己批自己」这个复核动作本身有价值，即使单人仓没有外部 reviewer。

### 根目录内容

根目录：`go.mod` / `README.md` / `.gitignore` / `AGENTS.md` / `GLOSSARY.md` / `.vscode/` / `docs/` / `src/` / `tools/`。

- `index.go`（空的 `package main`）已删除：根目录不该有 `main` 包。
- `README.md` 的 552 行原文不归档，正文由用户手写。原文可从 `eff15b9:README.md` 取回。

### CI

暂不加 GitHub Actions，留明确触发条件：第一次引入第三方依赖（Gin / GORM）时补 `.github/workflows/go.yml`，跑 `gofmt -l` / `go vet` / `go build`。

## 待定分支

- 自动创建目录的 CLI（`tools/`）的形状：生成什么骨架、编号规则、是否交互。用户会自己做，动之前需确认。

## 已否分支

- **学习内容放根级、不设 `src/`**（我推荐，用户否决）：原本的理由是 Go 的 `src/` 有 GOPATH 专指含义、导入路径会多一级、根目录只有 6 个条目并不挤。被否的理由是内容与工具需要分区，`src/` 提供这个边界。记录在此是为了下次不要重提同一个建议。
- **迁到 `apps/go-learning`**：它是学习档案不是应用；`playgrounds/rn-start` 已确立「playgrounds 下可以有真仓」的先例。
- **模块路径保持 `go-start`**：无法被 `go get`；仓库已 public，改用远端路径的成本只有一行。
- **改用 `CONTEXT.md` 迎合工作区默认**：丢掉了「术语表只放术语」这条边界。
- **每个练习独立 `go.mod` / 提前上 `go.work`**：多出成倍的 `go.mod`+`go.sum` 要维护，换来的是学习阶段用不到的多版本隔离。
- **强制每个知识点齐三份**：写不出作业的知识点会被迫造假内容。
- **作业配 `_test.go` 或批量运行脚本**：一次性答案不值得配测试。
- **取消 PR 流程、直接提交 `main`**（我推荐，用户否决）：单人仓无 CI 无协作者，PR 看似只是仪式。被否的理由是复核动作本身有价值。`AGENTS.md` 里已注明「别以单人仓没必要为由省掉」。
- **为 `src/`、单根模块、bootstrap 例外写 ADR**：`src/` 的理由已完整记在本文件与 `AGENTS.md`，ADR 收的是「难以逆转 + 不解释会困惑 + 真有权衡」的决策；这几项在 `AGENTS.md` 里足够自解释。若将来涉及导入路径的实际代价，再补 ADR。
- **把首个提交推成分支再手动合 PR**：空仓没有 base 分支，PR 开不出来。
- **`README.md` 原文移入 `docs/reference/` 存档**：那 552 行是一次性设计建议，结论已进本文件。
