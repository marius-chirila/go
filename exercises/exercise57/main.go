package main

import (
	"fmt"
	"runtime"
)

func main() {
	gr := 100
	for i := 0; i < gr; i++ {
		go func() {
			for i := 0; i < gr; i++ {
				c := factorial(i)
				d := consumeFact(c)
				for j := range d {
					fmt.Println("Each factorial:", j, "No Goroutine:", runtime.NumGoroutine())
				}
			}
		}()
	}

	// for n := range c {
	// 	fmt.Println(n)
	// }
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

func consumeFact(c chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range c {
			out <- i
		}
		close(out)
	}()
	return out
}
