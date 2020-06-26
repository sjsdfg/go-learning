package chap08

import "testing"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func TestLiteral(t *testing.T) {
	subscriber := Subscriber{
		Name: "",
		Rate: 0,
	}
	t.Log(subscriber)

	t.Log(subscriber.Active)
}
