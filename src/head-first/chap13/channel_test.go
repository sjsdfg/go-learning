package chap13

import (
	"fmt"
	"testing"
	"time"
)

func greeting(channel chan string) {
	channel <- "greeting"
}

func TestChannel(t *testing.T) {
	channel := make(chan string)
	go greeting(channel)

	fist := <-channel
	t.Log(fist)
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(time.Second)
	}
	fmt.Println(name, "weak up")
}

func send(channel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("*** sending value ***")
	channel <- "a"
	fmt.Println("*** sending value ***")
	channel <- "b"
}

func TestSending(t *testing.T) {
	channel := make(chan string)
	go send(channel)
	reportNap("receiving goroutine", 5)
	t.Log(<-channel)
	t.Log(<-channel)
}
