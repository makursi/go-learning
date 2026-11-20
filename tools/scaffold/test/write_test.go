package scaffoldtest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const handwrittenNote = "# 手写的内容，不该被覆盖\n"

func TestExistingNoteIsSkippedAndLeftAlone(t *testing.T) {
	root := fakeRepo(t)

	rel := "src/notes/01-basics/maps.md"
	abs := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatalf("造已有文件所在的目录：%v", err)
	}
	if err := os.WriteFile(abs, []byte(handwrittenNote), 0o644); err != nil {
		t.Fatalf("写已有笔记：%v", err)
	}

	stdout, stderr, code := runCLI(t, root, "note", "01-basics", "maps")

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
		t.Fatalf("读已有笔记：%v", err)
	}
	if string(got) != handwrittenNote {
		t.Errorf("已有文件被改动了：\n%s", got)
	}
}

// 题号是追加的：同名的题再跑一次会拿到下一个号，旧题目录原封不动。
func TestExistingExerciseIsNeverReused(t *testing.T) {
	root := fakeRepo(t)

	oldAbs := filepath.Join(root, "src", "homework", "02-concurrency", "channel", "01-hello", "main.go")
	if err := os.MkdirAll(filepath.Dir(oldAbs), 0o755); err != nil {
		t.Fatalf("造已有题目：%v", err)
	}
	const handwritten = "package main\n\n// 手写的题，不该被覆盖\nfunc main() {}\n"
	if err := os.WriteFile(oldAbs, []byte(handwritten), 0o644); err != nil {
		t.Fatalf("写已有题目：%v", err)
	}

	stdout, stderr, code := runCLI(t, root, "exercise", "02-concurrency", "channel", "hello")

	if code != 0 {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, 0, stderr)
	}
	if want := "create src/homework/02-concurrency/channel/02-hello/main.go"; !strings.Contains(stdout, want) {
		t.Errorf("stdout 应报告 %q，得到 %q", want, stdout)
	}

	got, err := os.ReadFile(oldAbs)
	if err != nil {
		t.Fatalf("读旧题目：%v", err)
	}
	if string(got) != handwritten {
		t.Errorf("旧题目被改动了：\n%s", got)
	}
}

func TestWriteFailureExitsOne(t *testing.T) {
	root := fakeRepo(t)

	homeworkDir := filepath.Join(root, "src", "homework")
	if err := os.MkdirAll(homeworkDir, 0o755); err != nil {
		t.Fatalf("造作业区：%v", err)
	}
	// 把主题目录的位置占成一个文件，建目录必然失败。
	if err := os.WriteFile(filepath.Join(homeworkDir, "02-concurrency"), []byte("占位"), 0o644); err != nil {
		t.Fatalf("占位文件：%v", err)
	}

	_, stderr, code := runCLI(t, root, "exercise", "02-concurrency", "channel", "hello")

	if code != 1 {
		t.Fatalf("退出码 = %d，想要 %d", code, 1)
	}
	if !strings.HasPrefix(stderr, "error:") {
		t.Errorf("stderr 应以 error: 开头，得到 %q", stderr)
	}
}
