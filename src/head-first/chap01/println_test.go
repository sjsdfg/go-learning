package chap01

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestPrintln(t *testing.T) {
	t.Log(math.Floor(2.75))
	t.Log(strings.Title("this is a title test "))
}

func TestUnicode(t *testing.T) {
	// 这里的输出会是 unicode 数字 97 ，而不是输出字符 'a'
	fmt.Print('a')
}
