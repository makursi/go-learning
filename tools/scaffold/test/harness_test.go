package scaffoldtest

import (
	"bytes"
	"go/ast"
	"os"
	"path/filepath"
	"testing"

	"github.com/makursi/go-learning/tools/scaffold/internal/scaffold"
)

// fakeRepo 造一个最小的假仓库根：只需要一份 module 路径正确的 go.mod。
func fakeRepo(t *testing.T) string {
	t.Helper()

	root := t.TempDir()
	gomod := "module " + "github.com/makursi/go-learning" + "\n\ngo 1.27.1\n"
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte(gomod), 0o644); err != nil {
		t.Fatalf("写 go.mod：%v", err)
	}

	return root
}

// runCLI 是测试唯一的入口：一条命令的输入是参数与工作目录，输出是两个流与退出码。
func runCLI(t *testing.T, cwd string, args ...string) (stdout, stderr string, code int) {
	t.Helper()

	var out, errOut bytes.Buffer
	code = scaffold.Run(args, cwd, &out, &errOut)

	return out.String(), errOut.String(), code
}

// hasMainFunc 判断文件里有没有无接收者的 main 函数。
func hasMainFunc(file *ast.File) bool {
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Recv == nil && fn.Name.Name == "main" {
			return true
		}
	}

	return false
}
