# hello-world


## **编写代码Hello World代码*
1. 进入用户主目录：
   - Linux/Mac：`cd`
   - Windows：`cd %HOMEPATH%`

2. 创建并进入项目目录：
   ```
   mkdir hello
   cd hello
   ```

3. 初始化模块（启用依赖管理）：
   ```
   go mod init example/hello
   ```
   会生成 `go.mod` 文件，记录模块路径和依赖。

4. 创建 `hello.go`，写入以下代码：
   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!")
   }
   ```
   - `package main`：声明主包
   - `import "fmt"`：导入标准库格式化包
   - `main` 函数：程序入口，默认执行

5. 运行代码：
   ```
   go run .
   ```
   输出：`Hello, World!`

   可用 `go help` 查看所有 go 命令。

## **调用外部包**
1. 在 [pkg.go.dev](https://pkg.go.dev) 搜索需要的包（示例搜索 “quote”）。
2. 选择 `rsc.io/quote` 模块的 `quote` 包，使用其中的 `Go` 函数。

3. 修改代码：
   ```go
   package main

   import "fmt"
   import "rsc.io/quote"

   func main() {
       fmt.Println(quote.Go())
   }
   ```

4. 整理依赖：
   ```
   go mod tidy
   ```
   自动下载所需模块（示例为 `rsc.io/quote v1.5.2`），并生成 `go.sum`。

5. 再次运行：
   ```
   go run .
   ```
   输出：`Don't communicate by sharing memory, share memory by communicating.`

**关键概念**
- 模块（module）：通过 `go.mod` 管理依赖
- 包（package）：同一目录下的文件组成，用于组织函数
- `go run`：编译并直接运行
- 外部依赖通过 `go mod tidy` 自动处理
