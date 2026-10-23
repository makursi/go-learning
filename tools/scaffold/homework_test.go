package main

import (
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHomeworkCreatesRunnableSkeleton(t *testing.T) {
	root := fakeRepo(t)

	stdout, stderr, code := runCLI(t, root, "homework", "01-basics", "fruit-count")

	if code != exitOK {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, exitOK, stderr)
	}

	rel := "src/homework/01-basics/fruit-count/main.go"
	if !strings.Contains(stdout, "create "+rel) {
		t.Errorf("stdout 应报告 create %s，得到 %q", rel, stdout)
	}

	content := readFile(t, root, rel)
	if content != mainGoTemplate {
		t.Errorf("作业骨架 = %q，想要与讲义同一份最小骨架", content)
	}
	if formatted, err := format.Source([]byte(content)); err != nil {
		t.Errorf("作业骨架不是合法 Go：%v", err)
	} else if string(formatted) != content {
		t.Errorf("作业骨架未经 gofmt：\n%s", content)
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
	root := fakeRepo(t)

	if _, _, code := runCLI(t, root, "homework", "01-basics", "fruit-count"); code != exitOK {
		t.Fatalf("跑作业生成失败，退出码 = %d", code)
	}

	for _, tree := range []string{"src/notes", "src/lectures"} {
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(tree))); !os.IsNotExist(err) {
			t.Errorf("%s 不该被作业生成建出来", tree)
		}
	}
}

func TestTitleFlagRejectedForKindsThatDoNotUseIt(t *testing.T) {
	root := fakeRepo(t)

	_, stderr, code := runCLI(t, root, "lecture", "--title", "演示", "01-basics", "hello")

	if code != exitUsage {
		t.Fatalf("退出码 = %d，想要 %d", code, exitUsage)
	}
	if !strings.HasPrefix(stderr, "error:") {
		t.Errorf("stderr 应以 error: 开头，得到 %q", stderr)
	}

	assertOnlyGoMod(t, root)
}
