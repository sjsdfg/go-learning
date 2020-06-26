package chap06

import (
	"os"
	"testing"
)

func TestCommandLineArgs(t *testing.T) {
	t.Log(os.Args)
	for i, arg := range os.Args {
		t.Log(i, arg)
	}
}
