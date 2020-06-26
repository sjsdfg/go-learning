package chap08

import "testing"

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

type subscriber struct {
	Name   string
	Rate   float64
	Active bool
	//
	Address
}

func TestEmbLiteral(t *testing.T) {
	address := Address{
		Street:     "",
		City:       "",
		State:      "123",
		PostalCode: "123",
	}

	employee := Employee{
		Name:        "",
		Salary:      0,
		HomeAddress: address,
	}

	t.Log(employee)

	sub := subscriber{
		Name:    "",
		Rate:    0,
		Active:  false,
		Address: Address{},
	}
	t.Logf("%#v", sub)
}
