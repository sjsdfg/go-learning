package chap02

import (
	"os"
	"testing"
)

func TestStats(t *testing.T) {
	stat, _ := os.Stat("D:\\workspace\\go\\go-learning\\Readme.md")
	t.Log(stat)
}
