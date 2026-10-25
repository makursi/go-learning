package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const modulePath = "github.com/makursi/go-learning"

// findRoot 从工作目录向上找本仓的根：go.mod 的 module 必须是 modulePath。
// 这样在 src/homework/... 里也能直接敲命令，而不必先 cd 回根目录。
func findRoot(cwd string) (string, error) {
	dir := cwd
	for {
		if data, err := os.ReadFile(filepath.Join(dir, "go.mod")); err == nil && moduleOf(data) == modulePath {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("不在 %s 仓库内：向上没有找到 module 为 %s 的 go.mod", cwd, modulePath)
		}

		dir = parent
	}
}

// moduleOf 取 go.mod 里的 module 路径；没有这一行返回空串。
func moduleOf(gomod []byte) string {
	for line := range strings.SplitSeq(string(gomod), "\n") {
		field, ok := strings.CutPrefix(strings.TrimSpace(line), "module")
		if ok {
			return strings.Trim(strings.TrimSpace(field), `"`)
		}
	}

	return ""
}
