package main

import (
	"fmt"
)

func main() {
	i := 20
	for i > -5 {
		fmt.Printf("Iterator is %d\t with type %T\t\n", i, i)
		i--
	}

	for {
		fmt.Printf("The iterator is %d\n", i)
		i++
		if i > 25 {
			break
		}
	}
}
