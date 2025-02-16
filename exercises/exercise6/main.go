package main

import (
	"fmt"
)

func main() {
	stringy := "Hello"
	inty := 42
	float := 3.14
	fmt.Printf("String %s and is of type %T\n", stringy, stringy)
	fmt.Printf("Int %d and is of type %T\n", inty, inty)
	fmt.Printf("Float %.2f and is of type %T\n", float, float)

}
