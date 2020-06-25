package chap05

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
	"testing"
)

func TestReadFile(t *testing.T) {
	result := list.New()
	file, err := os.Open("D:\\workspace\\go\\go-learning\\data\\float.txt")
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		float, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			t.Log(result)
			return
		}
		result.PushBack(float)
	}
	t.Logf("%v, %d", result, result.Len())
	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}
	if scanner.Err() != nil {
		t.Fatal(err)
	}

	for current := result.Front(); current != nil; current = current.Next() {
		t.Log(current)
	}
}
