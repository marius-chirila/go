package main

import (
	"fmt"
)

func main() {
	a, b, c := 747, 911, 90210
	fmt.Printf("Decimals: a - %d, b - %d, c - %d\n", a, b, c)
	fmt.Printf("Binaries: a - %b, b - %b, c - %b\n", a, b, c)
	fmt.Printf("Hexadecimals: a - %#X, b - %#X, c - %#X\n", a, b, c)

}
