package main

import (
	"fmt"
)

// var wg sync.WaitGroup

func main() {

	// wg.Add(2)
	c := gen()
	receive(c)
	// wg.Wait()
	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// wg.Done()
	return c

}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
	// wg.Done()
}
