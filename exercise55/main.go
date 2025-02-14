package main

import "fmt"

func main() {
	c := incrementor()

	// we don't need to range over that channel since there's only 1 value being sent to puller channel, the sum.
	i := <-puller(c)
	fmt.Println("Sum in channel incrementor is:", i)

}

// incrementor is adding to the channel all the values from 0->9.
func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

// puller gets values out of the channel that was return from incrementor and adds them and returns a channel with the sum of all the values that were in previous channel.
func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}
