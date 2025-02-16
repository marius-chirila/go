package main

import (
	"fmt"

	"math/rand"
)

func main() {

	for i := 0; i < 100; i++ { // runs previous exercise for 100 times
		x := rand.Intn(10)
		y := rand.Intn(10)
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

}
