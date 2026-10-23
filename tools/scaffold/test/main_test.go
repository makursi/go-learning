package scaffoldtest

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
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

func TestLectureCreatesRunnableSkeleton(t *testing.T) {
	root := fakeRepo(t)

	stdout, stderr, code := runCLI(t, root, "lecture", "01-basics", "hello")

	if code != 0 {
		t.Fatalf("退出码 = %d，想要 0；stderr = %q", code, stderr)
	}
	if stderr != "" {
		t.Errorf("stderr 应为空，得到 %q", stderr)
	}

	rel := "src/lectures/01-basics/hello/main.go"
	if !strings.Contains(stdout, "create "+rel) {
		t.Errorf("stdout 应报告 create %s，得到 %q", rel, stdout)
	}

	content, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatalf("读生成的入口文件：%v", err)
	}

	formatted, err := format.Source(content)
	if err != nil {
		t.Fatalf("生成的文件不是合法 Go：%v\n%s", err, content)
	}
	if !bytes.Equal(formatted, content) {
		t.Errorf("生成的文件未经 gofmt：\n%s", content)
	}

	file, err := parser.ParseFile(token.NewFileSet(), rel, content, 0)
	if err != nil {
		t.Fatalf("解析生成的文件：%v", err)
	}
	if file.Name.Name != "main" {
		t.Errorf("包名 = %q，想要 main", file.Name.Name)
	}
	if !hasMainFunc(file) {
		t.Errorf("生成的文件里没有 func main()：\n%s", content)
	}
}

func hasMainFunc(file *ast.File) bool {
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Recv == nil && fn.Name.Name == "main" {
			return true
		}
	}

	return false
}
