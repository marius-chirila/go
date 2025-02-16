package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where the initialization of the program occurs")
}
func main() {
	x := rand.Intn(251)
	fmt.Printf("The name of the variable is x and it's value is %d\n", x)
	/*
		Using the if conditional statement
	*/
	if x >= 0 && x <= 100 {
		fmt.Println("Between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("Between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("Between 201 and 250")
	}
	/*
		Using the switch logical statement
	*/
	y := rand.Intn(251)
	fmt.Printf("The name of the variable is y and it's value is %d\n", y)
	switch {
	case y <= 100:
		fmt.Println("Between 0 and 100")
	case y >= 101 && y <= 200:
		fmt.Println("Between 101 and 200")
	case y >= 201 && y <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("Out of range")
	}

}
