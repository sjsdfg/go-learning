package chap03

import "testing"

func TestPrintf(t *testing.T) {
	// 打印字符
	t.Logf("%v %v %v", "", "\t", "\n")
	// 原样打印
	// 这里的输出是 -> "" "\t" "\n"
	t.Logf("%#v %#v %#v", "", "\t", "\n")
}
