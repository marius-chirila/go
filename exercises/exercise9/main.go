package main

import (
	"fmt"
)

var stringy string = "Hello" // package variable

const inty int = 42 // package constant

func main() {
	float := 3.14 // short declaration operator
	fmt.Printf("String %s and is of type %T\n", stringy, stringy)
	fmt.Printf("Int %d and is of type %T\n", inty, inty)
	fmt.Printf("Float %.2f and is of type %T\n", float, float)
}
