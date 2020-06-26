package chap06

import "testing"

func TestCreate(t *testing.T) {
	// declare slice
	var sliceCreation []string
	// make slice
	sliceCreation = make([]string, 8)
	sliceCreation[0] = "t"

	t.Log(sliceCreation, len(sliceCreation), cap(sliceCreation))

	var literalsSlice = []int{18, 91, 7}
	t.Log(literalsSlice, len(literalsSlice), cap(literalsSlice))
}

func TestSliceOperator(t *testing.T) {
	var literalsSlice = []int{1, 2, 3, 4, 5}
	t.Log(literalsSlice, len(literalsSlice), cap(literalsSlice))

	slice := literalsSlice[1:3]
	t.Log(slice, len(slice), cap(slice))
}

func TestSliceAppend(t *testing.T) {
	array := []string{"a", "b", "c", "d", "e"}
	t.Log(array)

	slice := array[1:3]
	t.Log(array, slice)

	slice = append(slice, "x")
	t.Log(array, slice)

	slice = append(slice, "y")
	t.Log(array, slice)

	slice = append(slice, "z")
	t.Log(array, slice)
}
