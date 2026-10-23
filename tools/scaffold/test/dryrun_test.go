package scaffoldtest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDryRunPrintsPlanWithoutTouchingDisk(t *testing.T) {
	root := fakeRepo(t)

	stdout, stderr, code := runCLI(t, root, "note", "--dry-run", "01-basics", "maps")

	if code != 0 {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, 0, stderr)
	}

	rel := "src/notes/01-basics/maps.md"
	if !strings.Contains(stdout, "create "+rel) {
		t.Errorf("stdout 应报告将被创建的是 %s，得到 %q", rel, stdout)
	}

	if _, err := os.Stat(filepath.Join(root, "src")); !os.IsNotExist(err) {
		t.Errorf("--dry-run 不该在磁盘上留东西（src/ 被建出来了）")
	}
}

func TestDryRunSkipsExistingTargetWithoutFailing(t *testing.T) {
	root := fakeRepo(t)

	rel := "src/notes/01-basics/maps.md"
	abs := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatalf("造已有笔记：%v", err)
	}
	if err := os.WriteFile(abs, []byte("# 手写\n"), 0o644); err != nil {
		t.Fatalf("写已有笔记：%v", err)
	}

	stdout, stderr, code := runCLI(t, root, "note", "--dry-run", "01-basics", "maps")

	if code != 0 {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, 0, stderr)
	}
	if !strings.Contains(stdout, "skip "+rel) {
		t.Errorf("stdout 应报告 skip %s，得到 %q", rel, stdout)
	}

	if got := readFile(t, root, rel); got != "# 手写\n" {
		t.Errorf("--dry-run 不该改动文件，得到 %q", got)
	}
}
