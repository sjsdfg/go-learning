package chap01

import (
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	t.Log(reflect.TypeOf(42))
	t.Log(reflect.TypeOf(3.1415926))
	t.Log(reflect.TypeOf(true))
	t.Log(reflect.TypeOf("hello world"))
}
