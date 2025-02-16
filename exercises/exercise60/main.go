// deadlocked // not solved
package main

import (
	"fmt"
)

func main() {
	c1 := incrementor("1")
	c2 := incrementor("2")
	total := make(chan int)

	go func() {
		for i := range c1 {
			for j := range c2 {
				total <- i + j
			}
		}
		close(total)
	}()

	for k := range total {
		fmt.Println(k)
	}
}

func incrementor(s string) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("Process: "+s+" printing:", i)
			c <- i
		}
		close(c)
	}()
	return c
}
