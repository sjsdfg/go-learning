package chap12

import (
	"fmt"
	"testing"
)

type Refrigerator []string

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

func (r Refrigerator) open() {
	fmt.Println("opening Refrigerator")
}

func (r Refrigerator) close() {
	fmt.Println("closing Refrigerator")
}

func (r Refrigerator) findFood(food string) {
	r.open()
	defer r.close()
	if result := find(food, r); result {
		fmt.Println("find given food", food)
	} else {
		fmt.Println("don't find given food", food)
	}
}

func TestDefer(t *testing.T) {
	refrigerator := Refrigerator{"fish", "bread"}
	refrigerator.findFood("fish")

	refrigerator.findFood("milk")
}
