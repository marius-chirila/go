package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	d := make(chan int, 1)

	d <- 43

	fmt.Println(<-c)
	fmt.Println(<-d)
}
