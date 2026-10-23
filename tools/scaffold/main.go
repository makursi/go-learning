package main

import (
	"fmt"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: 取当前目录：%v\n", err)
		os.Exit(exitIO)
	}

	os.Exit(run(os.Args[1:], cwd, os.Stdout, os.Stderr))
}
