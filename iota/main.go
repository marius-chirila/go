package main

import "fmt"

const (
	_         = iota
	DecimalKB = 1 << iota
	DecimalMB = 1 << iota
	DecimalGB = 1 << iota
	DecimalTB = 1 << iota
	DecimalPB = 1 << iota
	DecimalEB = 1 << iota
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
	PB = 1 << (iota * 10)
	EB = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", DecimalKB, DecimalKB)
	fmt.Printf("%d \t %b\n", DecimalMB, DecimalMB)
	fmt.Printf("%d \t %b\n", DecimalGB, DecimalGB)
	fmt.Printf("%d \t %b\n", DecimalTB, DecimalTB)
	fmt.Printf("%d \t %b\n", DecimalPB, DecimalPB)
	fmt.Printf("%d \t %b\n", DecimalEB, DecimalEB)
	fmt.Printf("%d \t %b\n", KB, KB)
	fmt.Printf("%d \t %b\n", MB, MB)
	fmt.Printf("%d \t %b\n", GB, GB)
	fmt.Printf("%d \t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
}
