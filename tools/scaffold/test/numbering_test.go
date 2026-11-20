package scaffoldtest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTopicErrorSuggestsNextNumber(t *testing.T) {
	cases := []struct {
		desc   string
		topics []string
		want   string
	}{
		{"笔记区还空着时提示 01", nil, "01"},
		{"编号连续时提示下一个", []string{"01-basics", "02-functions"}, "03"},
		{"编号有空洞时取最大加一", []string{"01-basics", "03-functions"}, "04"},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			root := fakeRepo(t)
			for _, topic := range tc.topics {
				if err := os.MkdirAll(filepath.Join(root, "src", "notes", topic), 0o755); err != nil {
					t.Fatalf("造已有主题 %s：%v", topic, err)
				}
			}

			_, stderr, code := runCLI(t, root, "note", "functions", "maps")

			if code != 2 {
				t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, 2, stderr)
			}
			if !strings.Contains(stderr, "下一个可用编号") {
				t.Errorf("错误信息应给出下一个可用编号，得到 %q", stderr)
			}
			if !strings.Contains(stderr, tc.want) {
				t.Errorf("错误信息应包含 %s，得到 %q", tc.want, stderr)
			}
		})
	}
}

func TestTopicHintWhenTwoDigitNumbersRunOut(t *testing.T) {
	root := fakeRepo(t)
	if err := os.MkdirAll(filepath.Join(root, "src", "notes", "99-last"), 0o755); err != nil {
		t.Fatalf("造 99 号主题：%v", err)
	}

	_, stderr, code := runCLI(t, root, "note", "functions", "maps")

	if code != 2 {
		t.Fatalf("退出码 = %d，想要 %d", code, 2)
	}
	if strings.Contains(stderr, "100") {
		t.Errorf("不该提示 100——它过不了两位编号校验，得到 %q", stderr)
	}
	if !strings.Contains(stderr, "已用尽") {
		t.Errorf("应说明两位编号已用尽，得到 %q", stderr)
	}
}

func TestArgCountErrorNamesExpectedArgs(t *testing.T) {
	root := fakeRepo(t)

	_, stderr, code := runCLI(t, root, "exercise", "01-basics")

	if code != 2 {
		t.Fatalf("退出码 = %d，想要 %d", code, 2)
	}
	if !strings.Contains(stderr, "<topic>") || !strings.Contains(stderr, "<slug>") {
		t.Errorf("错误信息应点出期望的参数，得到 %q", stderr)
	}
}

func TestExerciseNumberTakesMaxOfExistingDirsPlusOne(t *testing.T) {
	root := fakeRepo(t)
	for _, name := range []string{"01-swap", "03-maps"} {
		if err := os.MkdirAll(filepath.Join(root, "src", "homework", "01-basics", name), 0o755); err != nil {
			t.Fatalf("造已有题目 %s：%v", name, err)
		}
	}

	stdout, stderr, code := runCLI(t, root, "exercise", "01-basics", "hello")

	if code != 0 {
		t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, 0, stderr)
	}
	if want := "create src/homework/01-basics/04-hello/main.go"; !strings.Contains(stdout, want) {
		t.Errorf("编号有空洞时取最大加一，stdout = %q，想要含 %q", stdout, want)
	}
}

func TestExerciseNumberExhausted(t *testing.T) {
	root := fakeRepo(t)
	if err := os.MkdirAll(filepath.Join(root, "src", "homework", "01-basics", "99-last"), 0o755); err != nil {
		t.Fatalf("造 99 号题目：%v", err)
	}

	_, stderr, code := runCLI(t, root, "exercise", "01-basics", "hello")

	if code != 2 {
		t.Fatalf("退出码 = %d，想要 %d", code, 2)
	}
	if !strings.Contains(stderr, "已用尽") {
		t.Errorf("应说明两位题号已用尽，得到 %q", stderr)
	}
	if _, err := os.Stat(filepath.Join(root, "src", "homework", "01-basics", "100-hello")); !os.IsNotExist(err) {
		t.Errorf("编号用尽时不该创建任何东西")
	}
}
