# 手动提交 PR

改之前看状态，改完跑门禁，然后按这七步走。别在 `main` 上直接提交。

## 1. 开工

```sh
git switch main
git pull --ff-only
git switch -c docs/note-topic     # 分支名取 <类型>/<短 slug>
```

## 2. 自查

```sh
git status                        # 看清加了什么、删了什么
git diff                          # 逐行看完再提交
```

改了 Go 代码就跑门禁：

```sh
gofmt -l .                        # 必须无输出
go vet ./...
go build ./...
go test ./...
```

## 3. 提交

```sh
git add <path>                    # 按路径加，别图省事用 git add -A
git commit -m "docs: 一句话说清做了什么"
```

提交信息写成 `<类型>: <中文说明>`。类型取 `feat` / `fix` / `docs` / `refactor` / `test` / `chore`。

## 4. 推送

```sh
git push -u origin docs/note-topic
```

## 5. 开 PR

先把正文写进临时文件，再创建：

```sh
gh pr create --base main --head docs/note-topic \
  --title "docs: 一句话说清做了什么" \
  --body-file pr-body.md
```

正文按这三段写：

```markdown
## Summary

改了什么。给一棵树、一段 diff 或一行命令，别复述代码。

## Evidence

怎么证明有效。贴真实输出：命令结果、测试通过、前后对比。

## Merge Danger

**Door:** two-way        # 能直接 revert 就是 two-way；动了数据或对外契约就是 one-way
**Blast Radius:** local  # 影响范围：local / repo-wide / 对外

说清回滚代价与影响面。
```

创建完删掉临时文件。

## 6. 合并

自查通过再合并，只 squash：

```sh
gh pr merge <编号> --squash --delete-branch
```

## 7. 收尾

```sh
git switch main
git pull --ff-only
git branch -d <分支名>            # 远端已被删，本地清掉
git remote prune origin           # 清掉过期的远端引用
```

## 硬规矩

- 不要在 `main` 上直接提交。
- 一次 PR 只做一件事。顺手改到的无关文件先放回去。
- 合并前看完整个 diff。
- 合并一律 `--squash`，让 `main` 保持一条线。
