# go-learning

个人Golang学习记录

## 文档

### 官方 · 参考

- [Go 语言规范](https://go.dev/ref/spec) — 语法、类型、求值的最终解释权
- [标准库索引](https://pkg.go.dev/std) — 全部标准库包的 API
- [Go 命令与工具](https://go.dev/doc/cmd) — `go` 命令及自带工具的完整用法
- [Go Modules 参考](https://go.dev/ref/mod) — 模块、版本、代理、校验的完整规则
- [go.mod 参考](https://go.dev/doc/modules/gomod-ref) — `go.mod` 每条指令的含义
- [内存模型](https://go.dev/ref/mem) — 什么情况下一个 goroutine 能看到另一个的写
- [版本发布历史](https://go.dev/doc/devel/release) — 每个版本改了什么
- [Go 博客](https://go.dev/blog) — 切片内部结构、反射三定律、defer 的实现都在这里
- [FAQ](https://go.dev/doc/faq) — 常见问题的官方回答
- [Go Wiki](https://go.dev/wiki) — 社区维护的官方 wiki

### 官方 · 学习

- [A Tour of Go](https://go.dev/tour/) — 四段式交互教程，含并发与练习
- [Effective Go](https://go.dev/doc/effective_go) — 怎么写出地道的 Go
- [如何写 Go 代码](https://go.dev/doc/code) — 包、模块、测试的最小工作流
- [教程：创建模块](https://go.dev/doc/tutorial/create-module) — 函数、错误、数组、map、单测一条龙
- [教程：多模块工作区](https://go.dev/doc/tutorial/workspaces) — `go.work` 到底解决什么问题
- [教程：泛型入门](https://go.dev/doc/tutorial/generics)
- [教程：模糊测试](https://go.dev/doc/tutorial/fuzz)
- [教程：用 Gin 写 REST API](https://go.dev/doc/tutorial/web-service-gin)
- [教程：访问关系型数据库](https://go.dev/doc/tutorial/database-access)
- [文档注释规范](https://go.dev/doc/comment) — `//` 注释如何生成 `go doc`
- [垃圾回收指南](https://go.dev/doc/gc-guide) — GC 怎么工作、怎么调
- [诊断工具](https://go.dev/doc/diagnostics) — 出问题该用哪个工具查
- [数据竞争检测](https://go.dev/doc/articles/race_detector) — `-race` 的手册
- [性能剖析](https://go.dev/blog/pprof) — pprof 的使用与读图
- [覆盖率](https://go.dev/doc/build-cover) — `-cover` 与覆盖率报告
- [PGO 入门](https://go.dev/doc/pgo) — 按运行剖面优化编译
- [Go by Example](https://gobyexample.com/) — 英文示例集，每个知识点一段可运行代码（第三方）

### 中文

- [Go 编程圣典](https://go-docs.pages.dev/) — 个人系统化中文教程，按主题成篇（非官方）
- [golang.google.cn](https://golang.google.cn/) — `go.dev` 的国内镜像，内容同源，仍是英文

### 生态 · 框架与库

- [Gin](https://gin-gonic.com/docs/) — 最常用的 Web 框架
- [Echo](https://echo.labstack.com/docs) — 路由与中间件风格更接近标准库
- [Fiber](https://docs.gofiber.io/) — 基于 fasthttp，吞吐优先
- [chi](https://github.com/go-chi/chi) — 纯 `net/http` 风格的路由器
- [GORM](https://gorm.io/docs/) — 最常用的 ORM
- [sqlc](https://docs.sqlc.dev/) — 由 SQL 生成类型安全的 Go 代码，不走反射
- [cobra](https://cobra.dev/) — 命令行应用骨架
- [viper](https://github.com/spf13/viper) — 配置读取（文件、环境变量、远程）
- [zap](https://github.com/uber-go/zap) — 结构化日志
- [testify](https://github.com/stretchr/testify) — 断言与 mock
- [wire](https://github.com/google/wire) — 编译期依赖注入

### 生态 · 云原生与基础设施

- [Docker](https://docs.docker.com/) — 容器与镜像构建，`Dockerfile` 与 `docker run` 的权威手册
- [Kubernetes](https://kubernetes.io/docs/home/) — 容器编排，用 Go 写的最大的开源系统之一
- [client-go](https://github.com/kubernetes/client-go) — 从 Go 里操作 K8s 的官方客户端，写 operator 的起点
- [Kubebuilder](https://book.kubebuilder.io/) — 用 Go 写自定义资源与 controller 的教程式手册
- [Helm](https://helm.sh/docs/) — K8s 包管理，本身用 Go 写
- [gRPC-Go](https://grpc.io/docs/languages/go/) — Go 的 gRPC 官方实现与入门
- [Protocol Buffers](https://protobuf.dev/) — gRPC 的序列化格式与 `.proto` 语言指南
- [Prometheus client_golang](https://github.com/prometheus/client_golang) — 在 Go 服务里暴露监控指标
- [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/) — 链路追踪与指标的统一埋点
- [etcd](https://etcd.io/docs/) — 分布式键值存储，Go 写的，读一致性实现的样本
- [NATS](https://docs.nats.io/) — 轻量消息系统，Go 写的
- [Caddy](https://caddyserver.com/docs/) — 自动 HTTPS 的反向代理，Go 写的
- [Traefik](https://doc.traefik.io/traefik/) — 面向容器的反向代理与网关
- [CNCF Landscape](https://landscape.cncf.io/) — 云原生全景图，找项目定位用

### 生态 · 用 Go 写的知名项目

- [containerd](https://containerd.io/docs/) — 容器运行时，Docker 底层就是它
- [Podman](https://podman.io/docs) — 无守护进程的容器引擎
- [CRI-O](https://cri-o.io/) — 专为 Kubernetes 的容器运行时
- [k3s](https://docs.k3s.io/) — 单二进制 Kubernetes
- [kind](https://kind.sigs.k8s.io/) — 用容器跑 K8s 集群，本地测试用
- [minikube](https://minikube.sigs.k8s.io/docs/) — 本地单机 K8s
- [Kustomize](https://kustomize.io/) — 不写模板的 K8s 配置层叠
- [Argo CD](https://argo-cd.readthedocs.io/) — GitOps 持续部署
- [Flux](https://fluxcd.io/flux/) — 另一套 GitOps，与 Argo CD 同代
- [Tekton](https://tekton.dev/docs/) — 用 CRD 描述 CI/CD 流水线
- [Harbor](https://goharbor.io/docs/) — 企业级镜像仓库
- [Istio](https://istio.io/latest/docs/) — 服务网格，控制面是 Go，数据面是 Envoy（C++）
- [Dapr](https://docs.dapr.io/) — 分布式应用运行时，把状态、发布订阅做成 sidecar
- [Crossplane](https://docs.crossplane.io/) — 把云资源当成 K8s 对象管理
- [Consul](https://developer.hashicorp.com/consul/docs) — 服务发现与配置
- [Vault](https://developer.hashicorp.com/vault/docs) — 密钥管理
- [Nomad](https://developer.hashicorp.com/nomad/docs) — 调度器，比 K8s 轻
- [Terraform](https://developer.hashicorp.com/terraform/docs) — 基础设施即代码
- [Packer](https://developer.hashicorp.com/packer/docs) — 镜像构建
- [TiDB](https://docs.pingcap.com/tidb/stable) — 分布式 SQL 数据库，TiDB 层是 Go（TiKV 是 Rust）
- [CockroachDB](https://www.cockroachlabs.com/docs/) — 分布式 SQL，全 Go
- [Vitess](https://vitess.io/docs/) — MySQL 分库分表中间件，YouTube 出身
- [MinIO](https://min.io/docs/minio/linux/index.html) — S3 兼容对象存储
- [restic](https://restic.readthedocs.io/) — 去重加密备份
- [Prometheus](https://prometheus.io/docs/) — 监控系统本体，不只是客户端库
- [Jaeger](https://www.jaegertracing.io/docs/) — 链路追踪
- [Loki](https://grafana.com/docs/loki/latest/) — 日志聚合，只索引标签
- [Grafana](https://grafana.com/docs/grafana/latest/) — 后端是 Go，前端是 TypeScript
- [VictoriaMetrics](https://docs.victoriametrics.com/) — 时序数据库，Prometheus 的替代
- [k6](https://grafana.com/docs/k6/latest/) — 压测工具，脚本用 JS，引擎用 Go
- [Gitea](https://docs.gitea.com/) — 自建 Git 服务
- [Hugo](https://gohugo.io/documentation/) — 静态站点生成器，`go-docs.pages.dev` 就是它生成的
- [Syncthing](https://docs.syncthing.net/) — 点对点文件同步
- [rclone](https://rclone.org/docs/) — 各家云存储的统一命令行
- [frp](https://gofrp.org/en/docs/) — 内网穿透
- [Temporal](https://docs.temporal.io/) — 持久化工作流引擎
- [Geth](https://geth.ethereum.org/docs) — 以太坊官方客户端

### 生态 · 工具链

- [gopls](https://go.dev/gopls) — 官方语言服务器，编辑器的补全与跳转全靠它
- [Delve](https://github.com/go-delve/delve) — 调试器，`dlv debug` 的底层
- [staticcheck](https://staticcheck.dev/) — 比 `go vet` 更狠的静态检查
- [golangci-lint](https://golangci-lint.run/) — 把多种 linter 合成一条命令
- [air](https://github.com/air-verse/air) — 改动自动重编译重启
- [Go Playground](https://go.dev/play) — 在线运行片段，贴代码问问题用它

### 生态 · 社区与资讯

- [Go 源码](https://github.com/golang/go) — 读标准库实现的地方
- [Go Issue Tracker](https://github.com/golang/go/issues) — 语言与工具链的问题都在这里
- [awesome-go](https://github.com/avelino/awesome-go) — 按领域分类的库清单
- [Golang Weekly](https://golangweekly.com/) — 每周生态动态
- [r/golang](https://www.reddit.com/r/golang/)
- [Gopher Slack](https://invite.slack.golangbridge.org/)
