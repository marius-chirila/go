package main

import "fmt"

const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

const (
	c3 = iota // c0 == 0
	c4
	c5
	c6
)

const (
	_         = iota
	DecimalKB = 1 << iota
	DecimalMB = 1 << iota
	DecimalGB = 1 << iota
	DecimalTB = 1 << iota
	DecimalPB = 1 << iota
	DecimalEB = 1 << iota
)

func main() {
	fmt.Println(c0, c1, c2)
	fmt.Println(c3, c4, c5, c6)
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", DecimalKB, DecimalKB)
	fmt.Printf("%d \t %b\n", DecimalMB, DecimalMB)
	fmt.Printf("%d \t %b\n", DecimalGB, DecimalGB)
	fmt.Printf("%d \t %b\n", DecimalTB, DecimalTB)
	fmt.Printf("%d \t %b\n", DecimalPB, DecimalPB)
	fmt.Printf("%d \t %b\n", DecimalEB, DecimalEB)
}
