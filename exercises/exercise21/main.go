package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("Outerloop %d\t Inner loop %d \n\t", i, j)
		}
	}
}
