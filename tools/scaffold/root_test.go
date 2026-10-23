package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunsFromAnyDirectoryInsideRepo(t *testing.T) {
	root := fakeRepo(t)

	sub := filepath.Join(root, "src", "lectures", "01-basics", "hello")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatalf("造的嵌套目录：%v", err)
	}

	stdout, stderr, code := runCLI(t, sub, "lecture", "02-functions", "demo")

	if code != exitOK {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, exitOK, stderr)
	}

	rel := "src/lectures/02-functions/demo/main.go"
	if !strings.Contains(stdout, "create "+rel) {
		t.Errorf("stdout 应报告 create %s，得到 %q", rel, stdout)
	}
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(rel))); err != nil {
		t.Errorf("文件应落在仓库根之下：%v", err)
	}
}

func TestRefusesOutsideRepo(t *testing.T) {
	outside := t.TempDir()

	_, stderr, code := runCLI(t, outside, "lecture", "01-basics", "hello")

	if code != exitUsage {
		t.Fatalf("退出码 = %d，想要 %d", code, exitUsage)
	}
	if !strings.HasPrefix(stderr, "error:") {
		t.Errorf("stderr 应以 error: 开头，得到 %q", stderr)
	}

	entries, err := os.ReadDir(outside)
	if err != nil {
		t.Fatalf("读仓库外目录：%v", err)
	}
	if len(entries) != 0 {
		t.Errorf("仓库外不该创建任何东西，得到 %d 项", len(entries))
	}
}
