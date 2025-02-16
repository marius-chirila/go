package main

import (
	"fmt"
)

func main() {
	in := gen()
	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			x := <-fact(n)
			out <- x
		}
		close(out)
	}()
	return out
}

func fact(n int) chan int { // fan in
	total := make(chan int)
	go func() {
		calc := 1
		for i := n; i > 0; i-- {
			calc *= i

		}
		total <- calc
		close(total)
	}()
	return total

}
