package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNoteDefaultsHeadingToName(t *testing.T) {
	root := fakeRepo(t)

	stdout, stderr, code := runCLI(t, root, "note", "01-basics", "maps")

	if code != exitOK {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, exitOK, stderr)
	}

	rel := "src/notes/01-basics/maps.md"
	if !strings.Contains(stdout, "create "+rel) {
		t.Errorf("stdout 应报告 create %s，得到 %q", rel, stdout)
	}

	content := readFile(t, root, rel)
	if content != "# maps\n\n" {
		t.Errorf("笔记内容 = %q，想要标题加一个空行", content)
	}
}

func TestNoteTitleFlagBecomesHeading(t *testing.T) {
	root := fakeRepo(t)

	stdout, stderr, code := runCLI(t, root, "note", "--title", "映射（Maps）", "01-basics", "maps")

	if code != exitOK {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, exitOK, stderr)
	}

	rel := "src/notes/01-basics/maps.md"
	if !strings.Contains(stdout, "create "+rel) {
		t.Errorf("stdout 应报告 create %s，得到 %q", rel, stdout)
	}

	content := readFile(t, root, rel)
	if content != "# 映射（Maps）\n\n" {
		t.Errorf("笔记内容 = %q，想要 --title 原样成为标题", content)
	}
}

func readFile(t *testing.T, root, rel string) string {
	t.Helper()

	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatalf("读 %s：%v", rel, err)
	}

	return string(data)
}
