package chap09

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string
	No   string
}

func (student *Student) toString() string {
	return fmt.Sprintf("name is %s, No is %s", student.Name, student.No)
}

func TestMethod(t *testing.T) {
	student := Student{
		Name: "123",
		No:   "345",
	}
	t.Log(student.toString())
}
