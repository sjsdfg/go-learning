package chap12

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRecover(t *testing.T) {
	defer reportPanic()
	panic(fmt.Errorf("test"))
}

func reportPanic() {
	result := recover()
	if result == nil {
		return
	}
	if msg, ok := result.(error); ok {
		fmt.Println(msg)
	} else {
		fmt.Println(reflect.TypeOf(msg))
	}
}
