package main

import "fmt"

func main() {
	x, y := 83/40, 83%40 // have a modulus result of 83 divided by 40 in float
	fmt.Println(float64(x) + 0.1*float64(y))
	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			fmt.Printf("This is not an even number %d\n", i)
		}
	}

}
