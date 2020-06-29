package chap13

import (
	"fmt"
	"testing"
	"time"
)

func print(value string) {
	for i := 0; i < 50; i++ {
		fmt.Print(value)
	}
}

func TestGoroutine(t *testing.T) {
	go print("a")
	go print("b")

	time.Sleep(time.Second)
	fmt.Println("ending")
}
