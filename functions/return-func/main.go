package main

import (
	"fmt"
)

func main() {
	y := bar()
	fmt.Println(y())
}

func bar() func() int {
	return func() int {
		return 42
	}
}
