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

// exerciseSkeleton 是工具生成的一道题的骨架：题面占位注释 + 最小可运行的 main。
const exerciseSkeleton = "// 题目：\n\npackage main\n\nfunc main() {\n}\n"

func TestExerciseCreatesRunnableSkeleton(t *testing.T) {
	root := fakeRepo(t)

	stdout, stderr, code := runCLI(t, root, "exercise", "02-concurrency", "channel", "swap")

	if code != 0 {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, 0, stderr)
	}
	if stderr != "" {
		t.Errorf("stderr 应为空，得到 %q", stderr)
	}

	rel := "src/homework/02-concurrency/channel/01-swap/main.go"
	if !strings.Contains(stdout, "create "+rel) {
		t.Errorf("stdout 应报告 create %s，得到 %q", rel, stdout)
	}

	content := readFile(t, root, rel)
	if content != exerciseSkeleton {
		t.Errorf("题目骨架 = %q，想要题面占位加最小可运行骨架", content)
	}

	formatted, err := format.Source([]byte(content))
	if err != nil {
		t.Fatalf("题目骨架不是合法 Go：%v\n%s", err, content)
	}
	if !bytes.Equal(formatted, []byte(content)) {
		t.Errorf("题目骨架未经 gofmt：\n%s", content)
	}

	file, err := parser.ParseFile(token.NewFileSet(), rel, []byte(content), parser.ParseComments)
	if err != nil {
		t.Fatalf("解析题目骨架：%v", err)
	}
	if file.Name.Name != "main" {
		t.Errorf("包名 = %q，想要 main", file.Name.Name)
	}
	if !hasMainFunc(file) {
		t.Errorf("题目骨架里没有 func main()：\n%s", content)
	}
	if len(file.Comments) == 0 {
		t.Fatalf("题目骨架应带 // 题目： 占位注释：\n%s", content)
	}
	if file.Comments[0].Pos() > file.Package {
		t.Errorf("题目注释应在 package 之前，实际在其后：\n%s", content)
	}

	entries, err := os.ReadDir(filepath.Join(root, "src", "homework", "02-concurrency", "channel", "01-swap"))
	if err != nil {
		t.Fatalf("读题目目录：%v", err)
	}
	if len(entries) != 1 || entries[0].Name() != "main.go" {
		names := make([]string, 0, len(entries))
		for _, entry := range entries {
			names = append(names, entry.Name())
		}
		t.Errorf("题目目录应只有 main.go（题面写在它的顶部），得到 %v", names)
	}
}

func TestExerciseNumberingAppendsToHomework(t *testing.T) {
	root := fakeRepo(t)

	if _, stderr, code := runCLI(t, root, "exercise", "02-concurrency", "channel", "swap"); code != 0 {
		t.Fatalf("第一题生成失败，退出码 = %d；stderr = %q", code, stderr)
	}

	stdout, stderr, code := runCLI(t, root, "exercise", "02-concurrency", "channel", "maps")
	if code != 0 {
		t.Fatalf("第二题生成失败，退出码 = %d；stderr = %q", code, stderr)
	}

	if want := "create src/homework/02-concurrency/channel/02-maps/main.go"; !strings.Contains(stdout, want) {
		t.Errorf("第二题应接着编号成 02，stdout = %q，想要含 %q", stdout, want)
	}
}

func TestEachKindTouchesOnlyItsOwnTree(t *testing.T) {
	cases := []struct {
		desc       string
		args       []string
		wantAbsent string
	}{
		{"题目不碰笔记区", []string{"exercise", "02-concurrency", "channel", "swap"}, "src/notes"},
		{"笔记不碰作业区", []string{"note", "01-basics", "maps"}, "src/homework"},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			root := fakeRepo(t)

			if _, _, code := runCLI(t, root, tc.args...); code != 0 {
				t.Fatalf("%s 生成失败，退出码 = %d", tc.args[0], code)
			}

			if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(tc.wantAbsent))); !os.IsNotExist(err) {
				t.Errorf("%s 不该被 %s 建出来", tc.wantAbsent, tc.args[0])
			}
		})
	}
}

func TestTitleFlagRejectedForKindsThatDoNotUseIt(t *testing.T) {
	root := fakeRepo(t)

	_, stderr, code := runCLI(t, root, "exercise", "--title", "演示", "02-concurrency", "channel", "hello")

	if code != 2 {
		t.Fatalf("退出码 = %d，想要 %d", code, 2)
	}
	if !strings.HasPrefix(stderr, "error:") {
		t.Errorf("stderr 应以 error: 开头，得到 %q", stderr)
	}

	assertOnlyGoMod(t, root)
}
