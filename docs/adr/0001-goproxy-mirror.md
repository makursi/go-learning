# 模块代理固定为 goproxy.cn

本机网络下 `proxy.golang.org` 不可达（实测 TCP 443 连接超时），而它是 Go 工具链的默认代理，导致任何 `go install` / `go get` 都会挂起后失败。因此把 `GOPROXY` 持久化为 `https://goproxy.cn,direct`，通过 `go env -w` 写入 `%APPDATA%\go\env`，对本机所有 Go 项目生效——而不是只在这一条命令上临时设置。`GOSUMDB` 保持默认：goproxy.cn 会代取 sumdb，模块校验依然有效，无需关闭校验换取速度。

## Considered Options

- **保持默认 `proxy.golang.org,direct`**：兜底的 `direct` 要求能直连 github.com 等源站，本机同样不稳定，所以"默认值 + 直连"这条路走不通。
- **只在单条命令上挂临时环境变量**：不写任何全局配置，但每个新终端都要重来一遍；更麻烦的是编辑器内部调用的 go 命令读不到它，装 gopls 失败时会误判成"gopls 有问题"。
- **同时设置 `GOSUMDB=off`**：能少一次网络往返，但放弃了模块校验。既然镜像已经能代取 sumdb，这个代价不值得付。
