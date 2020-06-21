package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	/*

	 */
	flag.StringVar(&name, "name", "flying Joe", "the user name")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s! \n", name)
}
