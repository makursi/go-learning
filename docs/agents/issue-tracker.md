# Issue tracker：GitHub

本仓的 issue 与 spec 都记在 GitHub Issues，一律用 `gh` CLI 操作。

## 约定

- **建 issue**：`gh issue create --title "..." --body "..."`。多行正文用 heredoc。
- **读 issue**：`gh issue view <number> --comments`，配合 `jq` 过滤评论，并取回标签。
- **列 issue**：`gh issue list --state open --json number,title,body,labels,comments --jq '[.[] | {number, title, body, labels: [.labels[].name], comments: [.comments[].body]}]'`，按需要加 `--label` 与 `--state`。
- **评论**：`gh issue comment <number> --body "..."`
- **加标签 / 去标签**：`gh issue edit <number> --add-label "..."` / `--remove-label "..."`
- **关闭**：`gh issue close <number> --comment "..."`

仓库从 `git remote -v` 推断；在 clone 目录里跑 `gh` 会自动识别。

## PR 是否作为 triage 入口

**PRs as a request surface: no.**（这一行是 `/triage` 读的开关。本仓若开始把外部 PR 当功能请求，改成 `yes`。）

设为 `yes` 之后，PR 走与 issue 相同的标签与状态，命令换成 `gh pr` 系列：

- **读 PR**：`gh pr view <number> --comments`，看改动用 `gh pr diff <number>`。
- **列待 triage 的外部 PR**：`gh pr list --state open --json number,title,body,labels,author,authorAssociation,comments`，只保留 `authorAssociation` 为 `CONTRIBUTOR`、`FIRST_TIME_CONTRIBUTOR`、`NONE` 的，丢掉 `OWNER`/`MEMBER`/`COLLABORATOR`。
- **评论 / 打标签 / 关闭**：`gh pr comment`、`gh pr edit --add-label`/`--remove-label`、`gh pr close`。

GitHub 的 issue 与 PR 共用一个编号空间，所以裸的 `#42` 可能是两者之一：先用 `gh pr view 42` 判，失败再退回 `gh issue view 42`。

## skill 说「发布到 issue tracker」时

建一个 GitHub issue。

## skill 说「取相关工单」时

跑 `gh issue view <number> --comments`。

## Wayfinding 操作

供 `/wayfinder` 使用。**map** 是一个 issue，**子 issue** 是它下面的工单。

- **Map**：一个打了 `wayfinder:map` 标签的 issue，正文放 Notes / Decisions-so-far / Fog。`gh issue create --label wayfinder:map`。
- **子工单**：作为 map 的 GitHub 子 issue 关联（用 `gh api` 调 sub-issues 端点）。子 issue 不可用时，把子工单写进 map 正文的任务列表，并在子 issue 正文开头写 `Part of #<map>`。标签为 `wayfinder:<type>`（`research` / `prototype` / `grilling` / `task`）。被认领后指派给执行者。
- **阻塞**：用 GitHub 原生 issue 依赖，这是唯一在 UI 上可见的表示。建一条边：`gh api --method POST repos/<owner>/<repo>/issues/<child>/dependencies/blocked_by -F issue_id=<blocker-db-id>`，其中 `<blocker-db-id>` 是阻塞者的数字**数据库 id**（`gh api repos/<owner>/<repo>/issues/<n> --jq .id`，不是 `#number`，也不是 `node_id`）。GitHub 用 `issue_dependencies_summary.blocked_by` 报告当前未关闭的阻塞数，这就是实时的闸门。依赖功能不可用时，退化为子 issue 正文开头的 `Blocked by: #<n>, #<n>` 行。所有阻塞者都关闭，工单才算解锁。
- **找前端**：列出 map 下未关闭的子工单（`gh issue list --state open`，限定在子 issue / 任务列表内），丢掉有未关闭阻塞（`issue_dependencies_summary.blocked_by > 0`，或 `Blocked by` 行里还有未关闭的）或已有指派者的；按 map 顺序取第一个。
- **认领**：`gh issue edit <n> --add-assignee @me`，这是本次会话的第一次写操作。
- **解决**：`gh issue comment <n> --body "<答案>"`，然后 `gh issue close <n>`，再把上下文指针（要点 + 链接）追加到 map 的 Decisions-so-far。
