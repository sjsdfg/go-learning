package chap11

import (
	"fmt"
	"net/mail"
	"testing"
)

func AcceptAnything(empty interface{}) {
	fmt.Println(empty)
}

func TestEmptyInterface(t *testing.T) {
	AcceptAnything(123)
	AcceptAnything(12.5)
	AcceptAnything(mail.Address{
		Name:    "123",
		Address: "456",
	})
}
