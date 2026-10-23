# 域文档

工程类 skill 在探索代码库之前，如何消费本仓的域文档。

本仓是单上下文，且**不建** `CONTEXT.md`：域文档入口以 `AGENTS.md`「域文档」一节为准。

## 动手探索之前先读

- 根目录的 **`GLOSSARY.md`**。若存在 **`GLOSSARY-MAP.md`**，它指向各上下文的 `GLOSSARY.md`，读与本主题相关的那些。
- **`docs/adr/`**：读与本区域相关的 ADR。多上下文仓里还要看 `src/<上下文>/docs/adr/` 下的上下文级决策。
- **`docs/agents/design-notes.md`**：本仓的设计推导——`src/` 为什么存在、脚手架为什么长这样、哪些分支被否以及理由。提结构性改动之前先读它。

这些文件不存在时**静默跳过**：不提醒缺失，也不建议现在就建。`/domain-modeling`（由 `/grill-with-docs` 与 `/improve-codebase-architecture` 调用）会在术语或决策真正定下来时按需创建。

## 文件结构

单上下文（绝大多数仓）：

```
/
├── GLOSSARY.md
├── docs/adr/
│   ├── 0001-goproxy-mirror.md
└── src/
```

多上下文（根目录有 `GLOSSARY-MAP.md`）：

```
/
├── GLOSSARY-MAP.md
├── docs/adr/                          ← 全仓级决策
└── src/
    ├── ordering/
    │   ├── GLOSSARY.md
    │   └── docs/adr/                  ← 上下文级决策
    └── billing/
        ├── GLOSSARY.md
        └── docs/adr/
```

## 用术语表里的词

输出里出现领域概念时（issue 标题、重构提案、假设、测试名），用 `GLOSSARY.md` 定义的那个词，不要漂到它明确 `_Avoid_` 的同义词。

需要但术语表里还没有的概念，是一个信号：要么你在发明项目不用的语言（重新考虑），要么确实存在缺口（记下来交给 `/domain-modeling`）。

## ADR 冲突要挑明

输出与既有 ADR 矛盾时，明确说出来，不要悄悄覆盖：

> _与 ADR-0001（goproxy.cn）冲突，但值得重开，因为……_
