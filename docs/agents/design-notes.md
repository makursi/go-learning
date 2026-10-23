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

命名规则：`src/notes/02-data-structures/maps.md` ↔ `src/lectures/02-data-structures/maps/` 同名；作业按题名走 `src/homework/02-data-structures/fruit-count/`，不要求和讲义同名。

曾经靠 `src/notes/README.md` 的主题索引承载「缺的留空」。2026-10-23 用户删除了该文件并决定暂不重建（见文末「第二轮」），所以三边对齐现在的唯一载体是**同名约定 + 全文搜索**；索引已从「前提」降级为「待补的便利层」，不再是一条成立条件。

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

无。原「自动创建目录的 CLI」一项已结案，见文末「第二轮」。

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

## 第二轮：工具区 CLI（`scaffold`）

### 问题

在 `tools/` 里做一个 Go CLI，按学习内容的结构建出目录与骨架文件，与 `src/` 的内容隔离。用户给的三条要求：最小化、可拓展性强、可维护性高。

起点（2026-10-23 核实）：`tools/` 只有占位 `README.md`，全仓零 Go 代码；`src/` 下三个目录全空；`src/notes/README.md`（主题索引）已在工作区删除且未提交；单根模块 `github.com/makursi/go-learning`，go 1.27.1，`GOBIN` 在 PATH 上。

### 已定分支

**位置与代码形状**：`tools/scaffold/`，单个 `package main`，零第三方依赖。结构的唯一权威是 `kinds.go` 里的一张 `Kind` 表（种类 → 根目录 → 落点形态 → 内容模板 → 是否接受 `--title`）；规划层是纯函数（输入请求，输出待建的目录/文件清单），文件系统操作只在边缘。加一个种类 = 表里加一行 + 一个模板常量，判定逻辑一行不动。

不选 `main.go` + `internal/` 分层：同包测试足够用，多一层目录换不到任何边界。不选外部配置驱动的通用引擎：要额外养配置格式、schema 与错误处理，而本仓没有「别人来加种类」的诉求。不选 cobra：为三条命令引入首个第三方依赖，还会触发 `AGENTS.md` 里「第一次引依赖就补 CI」的条件。

**原子单元：单边**。一次调用只落一个种类的一个条目。三个种类：`note`（笔记文件）、`lecture`（讲义目录 + `main.go`）、`homework`（作业目录 + `main.go`）。主题目录按 `mkdir -p` 语义按需创建，**只建该种类用到的那棵树**——建讲义不会顺手在 `src/notes/` 下造一个空目录。不做「一次建齐三边」，因为它与两条既有决策冲突：允许缺位、作业名不必等于讲义名。将来真要「三边一起」，是加 `--with` 的增量特性，不返工。

**命令面**：

```sh
scaffold note     03-functions maps  --title "映射（Maps）"
scaffold lecture  03-functions maps
scaffold homework 03-functions fruit-count
```

每个种类一个独立的 `flag.FlagSet`，未知 flag 由标准库报错，不手写校验。不做 `topic` 种类（空目录 git 不跟踪，单边模型下主题目录是副作用，建出来立刻不可见）。暂不预留 `project` 种类（`src/projects/` 还不存在，加它是纯增量）。

**路径与编号**：主题必须带两位编号，`^\d{2}-[a-z0-9]+(-[a-z0-9]+)*$`；知识点与作业名用 `^[a-z0-9]+(-[a-z0-9]+)*$`。落点：`src/notes/<topic>/<kp>.md`、`src/lectures/<topic>/<kp>/main.go`、`src/homework/<topic>/<name>/main.go`。

编号**显式**：不接受「只给名字让工具分配」。自动编号会让「谁先建谁占号」变成隐藏状态。校验失败时错误信息附上「下一个可用编号是 03」（扫 `src/notes/` 现有最大编号 +1，格式 `%02d`），便利性靠提示而不是靠魔法。

**模板**：Go 常量，不进 `embed`、不用 `text/template`（模板目前没有第二个变量，常量比「模板名 → 文件」的弱链接更原子）。笔记 = `# <title>` 加一个空行，`--title` 缺省时用 slug；讲义与作业 = `package main` 加空 `main()`，tab 缩进，保证 `gofmt -l` 无输出。三类都不预置「要点」「TODO」这类逼人填空的段落。作业目录不放 `README.md`，题面按既有约定记在对应 Notes 里。

**运行时契约**：目录已存在则静默跳过；文件用 `O_EXCL` 创建，已存在则打 `skip` 继续。不提供 `--force`——那是学习内容的误删入口，要覆盖就手删。`--dry-run` 打印计划、不落盘。输出每行 `create|skip <仓库相对路径>`，结尾一行 `N created, M skipped`；错误走 stderr 并加 `error:` 前缀。退出码：`0` 全部成功或全部跳过、`1` I/O 失败、`2` 用法或校验失败（参数个数、slug、编号、不在本仓内）。

**仓库根发现**：从 cwd 向上找 `go.mod`，模块路径必须是 `github.com/makursi/go-learning`，否则报错退出 `2`。这样在 `src/lectures/...` 里敲命令也能用。`--root` 这次不做。

**调用方式**：日常 `go run ./tools/scaffold ...`；`tools/scaffold/README.md` 另外给一条 `go install ./tools/scaffold`（装进 `GOBIN`，已在 PATH 上），装完任何 cwd 都能敲 `scaffold`。

**测试**：要测。工具的价值正是「路径和编号不会错」，这恰恰是人手 `go run` 验不出来的。测试只从**一条 seam**进入——CLI 边界上的「参数 + 工作目录 + 两个输出流 → 退出码」，不引文件系统抽象、不 mock，在 `t.TempDir()` 里造带 `go.mod` 的假根用真文件系统断言。用例全是表驱动：三类落点与内容、slug 与编号拒绝（含「下一个可用编号」提示）、`--dry-run` 不落盘、幂等与 `skip`、只建该种类的树、仓库外拒绝、退出码 `0/1/2`。额外一条断言：生成的 `main.go` 过 `format.Source` 不变，骨架的 gofmt 干净不靠人记。

spec 里一度写过「slug / 规划 / 端到端三个测试文件」，落到了「一条 seam」上：把私有函数分别固定住会换来更准的失败定位，代价是每次重构都要改测试；工具只有几百行，前者不值得。

**索引**：工具不参与。用户删除 `src/notes/README.md` 后决定暂不重建，`AGENTS.md` 里「新增知识点时同时更新索引」那条一并删掉。三边对齐改由同名约定 + 全文搜索承载。

**文档与交付**：`tools/README.md` 从占位改写成工具区索引（链到 `tools/scaffold/README.md`）；`AGENTS.md` 更新三处——`tools/` 段落、门禁里的「空模块例外」（有包之后不再触发）、索引那条；`GLOSSARY.md` 增 `Scaffold` 并改 `Tool` 的定义。不写 ADR：三条门槛里的「难以逆转」不成立，工具随时可改。

### 已否分支

- **一次建齐三边 / 主题包**：与「允许缺位」「作业名不必等于讲义名」冲突，等于逼着为每个笔记造一个作业目录。
- **`topic` 种类**：空目录 git 不跟踪，单边模型下主题目录只是副作用。
- **`main.go` + `internal/` 分层**：同包测试够用，多一层目录无收益。
- **外部配置驱动的通用生成引擎**：配置格式、schema、错误处理都要养，学习仓不需要。
- **cobra**：首个第三方依赖，会触发 CI 条件，且只为三条命令。
- **索引派生 / 追加 / 只打印**：用户明确暂时不要索引；任何自动改 markdown 的方案都是在替一个不存在的文件做设计。
- **`--force` 覆盖**：学习内容的误删入口。
- **默认 dry-run、`--apply` 才写**：每次都要多敲一个 flag。
- **模板放 `embed` / `text/template`**：模板不含第二个变量，常量更原子。
- **交互向导**：标准库没有 prompt，不可脚本化，测试面变大。
- **自动分配编号**（只传名字）：隐藏状态，编号由先建者占。
- **为这个工具写 ADR**：三条门槛里的「难以逆转」不成立。

### 待定分支

无。
