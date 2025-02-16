package main

import (
	"fmt"
)

func main() {
	var a int8 = 127
	var b uint8 = 255
	var c int16 = 32767
	var d uint16 = 64535
	fmt.Printf("Biggest int8: %d and largest uint8: %d\n", a, b)
	fmt.Printf("Biggest int16: %d and largest uint16: %d\n", c, d)

}
