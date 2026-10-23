package main

import (
	"strings"
	"testing"
)

func TestUsageErrors(t *testing.T) {
	cases := []struct {
		desc string
		args []string
	}{
		{"没有参数", nil},
		{"只有种类", []string{"lecture"}},
		{"缺名字", []string{"lecture", "01-basics"}},
		{"未知种类", []string{"topic", "01-basics", "hello"}},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			root := fakeRepo(t)

			_, stderr, code := runCLI(t, root, tc.args...)

			if code != exitUsage {
				t.Fatalf("退出码 = %d，想要 %d；stderr = %q", code, exitUsage, stderr)
			}
			if !strings.HasPrefix(stderr, "error:") {
				t.Errorf("stderr 应以 error: 开头，得到 %q", stderr)
			}

			assertOnlyGoMod(t, root)
		})
	}
}
