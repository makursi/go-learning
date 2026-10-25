package scaffoldtest

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHomeworkCreatesRunnableSkeleton(t *testing.T) {
	root := fakeRepo(t)

	stdout, stderr, code := runCLI(t, root, "homework", "01-basics", "fruit-count")

	if code != 0 {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, 0, stderr)
	}
	if stderr != "" {
		t.Errorf("stderr 应为空，得到 %q", stderr)
	}

	rel := "src/homework/01-basics/fruit-count/main.go"
	if !strings.Contains(stdout, "create "+rel) {
		t.Errorf("stdout 应报告 create %s，得到 %q", rel, stdout)
	}

	content := readFile(t, root, rel)
	const want = "package main\n\nfunc main() {\n}\n"
	if content != want {
		t.Errorf("作业骨架 = %q，想要最小可运行骨架", content)
	}

	formatted, err := format.Source([]byte(content))
	if err != nil {
		t.Fatalf("作业骨架不是合法 Go：%v\n%s", err, content)
	}
	if !bytes.Equal(formatted, []byte(content)) {
		t.Errorf("作业骨架未经 gofmt：\n%s", content)
	}

	file, err := parser.ParseFile(token.NewFileSet(), rel, []byte(content), 0)
	if err != nil {
		t.Fatalf("解析作业骨架：%v", err)
	}
	if file.Name.Name != "main" {
		t.Errorf("包名 = %q，想要 main", file.Name.Name)
	}
	if !hasMainFunc(file) {
		t.Errorf("作业骨架里没有 func main()：\n%s", content)
	}

	entries, err := os.ReadDir(filepath.Join(root, "src", "homework", "01-basics", "fruit-count"))
	if err != nil {
		t.Fatalf("读作业目录：%v", err)
	}
	if len(entries) != 1 || entries[0].Name() != "main.go" {
		names := make([]string, 0, len(entries))
		for _, entry := range entries {
			names = append(names, entry.Name())
		}
		t.Errorf("作业目录应只有 main.go（题面记在笔记里），得到 %v", names)
	}
}

func TestEachKindTouchesOnlyItsOwnTree(t *testing.T) {
	cases := []struct {
		desc  string
		args  []string
		trees []string
	}{
		{"作业不碰笔记区", []string{"homework", "01-basics", "fruit-count"}, []string{"src/notes"}},
		{"笔记不碰作业区", []string{"note", "01-basics", "maps"}, []string{"src/homework"}},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			root := fakeRepo(t)

			if _, _, code := runCLI(t, root, tc.args...); code != 0 {
				t.Fatalf("%s 生成失败，退出码 = %d", tc.args[0], code)
			}

			for _, tree := range tc.trees {
				if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(tree))); !os.IsNotExist(err) {
					t.Errorf("%s 不该被 %s 建出来", tree, tc.args[0])
				}
			}
		})
	}
}

func TestTitleFlagRejectedForKindsThatDoNotUseIt(t *testing.T) {
	root := fakeRepo(t)

	_, stderr, code := runCLI(t, root, "homework", "--title", "演示", "01-basics", "hello")

	if code != 2 {
		t.Fatalf("退出码 = %d，想要 %d", code, 2)
	}
	if !strings.HasPrefix(stderr, "error:") {
		t.Errorf("stderr 应以 error: 开头，得到 %q", stderr)
	}

	assertOnlyGoMod(t, root)
}
