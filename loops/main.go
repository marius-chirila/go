package main

import (
	"fmt"
)

func main() {
	x := 42
	y := 5
	for i := 0; i < y; i++ {
		x += i
		fmt.Println(x)
	}
	fmt.Println("Final value of x is", x)

	for y >= 5 {
		fmt.Println(y)
		y++
		if y > 30 {
			break
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println("________________")
		for j := 0; j < 3; j++ {
			fmt.Printf("Outer loop %v Inner loop %v\n", i, j)
		}
	}
}
