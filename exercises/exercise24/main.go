package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := 0
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			c++
			fmt.Printf("%v \t %v\n", i, x)
		}
	}
	fmt.Printf("Total number of counts is: %v\n", c)
	fmt.Println("*****************************")
}
