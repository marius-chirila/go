package main

import (
	"fmt"
)

func main() {
	fmt.Print(`
	Showing up different ways to print values,
	like decimal, hexidecimal, binary etc.`)
	adams := 42
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Printf("Adams as a decimal: %d\n", adams)
	fmt.Printf("Adams as a hexideciaml: %x \n", adams)
	fmt.Printf("a,b,c,d,e,f: %#x %#x %#x %#x %#x %#x\n", a, b, c, d, e, f)
	fmt.Printf("a,b,c,d,e,f: %b %b %b %b %b %b\n", a, b, c, d, e, f)
}
