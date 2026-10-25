package scaffoldtest

import (
	"strings"
	"testing"
)

func TestRejectsInvalidTopicAndName(t *testing.T) {
	cases := []struct {
		desc  string
		topic string
		name  string
	}{
		{"主题只有一位编号", "3-basics", "hello"},
		{"主题编号后没有 slug", "03-", "hello"},
		{"主题含大写", "03-Basics", "hello"},
		{"主题缺编号", "basics", "hello"},
		{"主题含下划线", "03-basics_x", "hello"},
		{"名字含大写", "01-basics", "Hello"},
		{"名字含下划线", "01-basics", "hello_world"},
		{"名字含空格", "01-basics", "hello world"},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			root := fakeRepo(t)

			_, stderr, code := runCLI(t, root, "homework", tc.topic, tc.name)

			if code != 2 {
				t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, 2, stderr)
			}
			if !strings.HasPrefix(stderr, "error:") {
				t.Errorf("stderr 应以 error: 开头，得到 %q", stderr)
			}
			if !strings.Contains(stderr, tc.topic) && !strings.Contains(stderr, tc.name) {
				t.Errorf("stderr 应点出不合规的输入，得到 %q", stderr)
			}

			assertOnlyGoMod(t, root)
		})
	}
}
