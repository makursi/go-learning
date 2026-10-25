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

	_, stderr, code := runCLI(t, root, "homework", "01-basics")

	if code != 2 {
		t.Fatalf("退出码 = %d，想要 %d", code, 2)
	}
	if !strings.Contains(stderr, "<topic>") || !strings.Contains(stderr, "<name>") {
		t.Errorf("错误信息应点出期望的参数，得到 %q", stderr)
	}
}
