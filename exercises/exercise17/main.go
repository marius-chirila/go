package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y := rand.Intn(10), rand.Intn(10)
	fmt.Printf("The value of x is %d and the value of y is %d\n", x, y)
	if x < 4 && y < 4 {
		fmt.Println("Both x and y are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("Both x and y are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is between 4 and 6 inclusive")
	} else if y != 5 {
		fmt.Println("y is not equal to 5")
	} else {
		fmt.Println("None of the above cases were met!")
	}

	switch {
	case x < 4 && y < 4:
		fmt.Println("Both x and y are less than 4")
	case x > 6 && y > 6:
		fmt.Println("Both x and y are greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is between 4 and 6")
	case y != 5:
		fmt.Println("y is not equal to 5")
	default:
		fmt.Println("None of the above cases were met!")
	}
}
