package main

import (
	"strings"
	"testing"
)

func TestSummaryReportsCreatedThenSkipped(t *testing.T) {
	root := fakeRepo(t)

	first, stderr, code := runCLI(t, root, "note", "01-basics", "maps")
	if code != exitOK {
		t.Fatalf("第一次退出码 = %d；stderr = %q", code, stderr)
	}
	if !strings.Contains(first, "1 created, 0 skipped") {
		t.Errorf("第一次应汇总为 1 created, 0 skipped，得到 %q", first)
	}

	second, stderr, code := runCLI(t, root, "note", "01-basics", "maps")
	if code != exitOK {
		t.Fatalf("第二次退出码 = %d；stderr = %q", code, stderr)
	}
	if !strings.Contains(second, "0 created, 1 skipped") {
		t.Errorf("第二次应汇总为 0 created, 1 skipped，得到 %q", second)
	}
}

func TestDryRunSummaryCountsThePlannedAction(t *testing.T) {
	root := fakeRepo(t)

	stdout, stderr, code := runCLI(t, root, "homework", "--dry-run", "01-basics", "fruit-count")

	if code != exitOK {
		t.Fatalf("退出码 = %d；stderr = %q", code, stderr)
	}
	if !strings.Contains(stdout, "1 created, 0 skipped") {
		t.Errorf("--dry-run 的汇总应是计划值 1 created, 0 skipped，得到 %q", stdout)
	}
}
