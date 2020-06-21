package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = flag.String("name", "flying Joe", "test")
	flag.Parse()
	fmt.Printf("hello world, %s", *name)
}
