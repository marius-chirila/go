package main

import "fmt"

func main() {
	c := make(chan int)
	// send
	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

	// receive
	// bar(c)

	fmt.Println("About to exit")

}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// func bar(c <-chan int) {
// 	fmt.Println(<-c)
// }
