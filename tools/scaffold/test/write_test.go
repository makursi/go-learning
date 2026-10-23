package scaffoldtest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const handwritten = "package main\n\n// 手写的内容，不该被覆盖\nfunc main() {}\n"

func TestExistingFileIsSkippedAndLeftAlone(t *testing.T) {
	root := fakeRepo(t)

	rel := "src/lectures/01-basics/hello/main.go"
	abs := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatalf("造已有文件所在的目录：%v", err)
	}
	if err := os.WriteFile(abs, []byte(handwritten), 0o644); err != nil {
		t.Fatalf("写已有文件：%v", err)
	}

	stdout, stderr, code := runCLI(t, root, "lecture", "01-basics", "hello")

	if code != 0 {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, 0, stderr)
	}
	if !strings.Contains(stdout, "skip "+rel) {
		t.Errorf("stdout 应报告 skip %s，得到 %q", rel, stdout)
	}
	if strings.Contains(stdout, "create "+rel) {
		t.Errorf("已有的文件不该被报告成 create，得到 %q", stdout)
	}

	got, err := os.ReadFile(abs)
	if err != nil {
		t.Fatalf("读已有文件：%v", err)
	}
	if string(got) != handwritten {
		t.Errorf("已有文件被改动了：\n%s", got)
	}
}

func TestWriteFailureExitsOne(t *testing.T) {
	root := fakeRepo(t)

	lectures := filepath.Join(root, "src", "lectures")
	if err := os.MkdirAll(lectures, 0o755); err != nil {
		t.Fatalf("造讲义区：%v", err)
	}
	// 把主题目录的位置占成一个文件，建目录必然失败。
	if err := os.WriteFile(filepath.Join(lectures, "01-basics"), []byte("占位"), 0o644); err != nil {
		t.Fatalf("占位文件：%v", err)
	}

	_, stderr, code := runCLI(t, root, "lecture", "01-basics", "hello")

	if code != 1 {
		t.Fatalf("退出码 = %d，想要 %d", code, 1)
	}
	if !strings.HasPrefix(stderr, "error:") {
		t.Errorf("stderr 应以 error: 开头，得到 %q", stderr)
	}
}
