package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1 // this works as it's launched as a separate goroutine, doesn't block the println
	}()
	// c <- 1 // this blocks as it is in the same goroutine
	fmt.Println(<-c)
}
