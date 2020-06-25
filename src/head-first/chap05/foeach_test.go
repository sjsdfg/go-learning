package chap05

import (
	"testing"
)

func TestForeach(t *testing.T) {
	var arr = [3]int{1, 2, 3}
	t.Log(arr)
	t.Logf("%#v", arr)

	var sum = 0
	for _, val := range arr {
		t.Log(val)
		sum += val
	}
	t.Log(sum)
}
