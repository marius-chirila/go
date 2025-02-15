package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // need to close channel if we range over it
	}()
	fmt.Println(<-c)   // this pulls only 1 value from channel "c"
	for i := range c { //ranges over all the values in channel c
		fmt.Println(i)
	}
}
