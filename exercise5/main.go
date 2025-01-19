package main

import (
	"fmt"
)

var zeroValue int
var p *int

const io = iota

func main() {
	a, b := "Happy", "Birthday"
	anotherZeroValue := ""
	fmt.Printf("Print zero value of pointer: %v\n", p)
	fmt.Printf("%s %s\n", a, b)
	fmt.Printf("Print iota value: %d\n", io)
	fmt.Printf("Another zero value: %s\n", anotherZeroValue)
	anotherZeroValue = "hiya"
	fmt.Printf("Zero value: %d\n", zeroValue)
	fmt.Printf("Another zero value: %s\n", anotherZeroValue)

}
