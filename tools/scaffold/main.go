package main

import (
	"os"

	"github.com/makursi/go-learning/tools/scaffold/internal/scaffold"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		os.Stderr.WriteString("error: 取当前目录：" + err.Error() + "\n")
		os.Exit(1) // 与 scaffold 的「写文件系统失败」同一个退出码
	}

	os.Exit(scaffold.Run(os.Args[1:], cwd, os.Stdout, os.Stderr))
}
