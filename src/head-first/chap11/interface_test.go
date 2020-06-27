package chap11

import (
	"fmt"
	"testing"
)

type Interface interface {
	MethodWithoutParams()

	MethodWithParams(float64)

	MethodWithReturnValues() string
}

type CustomType int

func (custom CustomType) MethodWithoutParams() {
	fmt.Println("MethodWithoutParams")
}

func (custom CustomType) MethodWithParams(value float64) {
	fmt.Printf("MethodWithoutParams, values is %f\n", value)
}

func (custom CustomType) MethodWithReturnValues() string {
	return fmt.Sprintf("%#v", custom)
}

func TestInterface(t *testing.T) {
	var value Interface = CustomType(5)
	value.MethodWithoutParams()
	value.MethodWithParams(1.5)
	t.Log(value.MethodWithReturnValues())
}
