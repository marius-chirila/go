package main

import (
	"fmt"
)

func main() {
	var s int
	fmt.Println("Give an int")
	fmt.Scan(&s)
	if s < 42 {
		fmt.Println("Smaller than life!")
	} else if s == 42 {
		fmt.Println("The answer to life, the universe and everything!")
	} else if s > 42 {
		fmt.Println("Bigger than life!")
	} else {
		fmt.Println("Not an intiger")
	}

	switch {
	case s == 42: // if s == 42
		fmt.Println("The answer to life, the universe and everything!")
	case s < 42:
		fmt.Println("Smaller than life!")
	case s > 42:
		fmt.Println("Bigger than life!")
	default:
		fmt.Println("Not an intiger")
	}
}
