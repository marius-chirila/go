package main

import (
	"fmt"
)

func main() {
	// c := make(chan int)

	// c <- 42    This does not work, deadlock.

	// fmt.Println(<-c)
	c := make(chan int)
	go func() { // this works because it's a separate routine
		c <- 42
	}()
	fmt.Println(<-c)

	d := make(chan int, 1) // using channel buffer by the size of the second argument
	d <- 42                // only 1 value stored in the channel buffer as per size
	fmt.Println(<-d)

}
