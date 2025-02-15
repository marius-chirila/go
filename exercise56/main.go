package main

import "fmt"

var fact = make(chan int)

func main() {
	go factorial(4)
	fmt.Println("Total:", <-fact)
}

func factorial(n int) chan int {
	calc := 1
	for i := n; i > 0; i-- {
		calc *= i
	}
	fact <- calc
	close(fact)
	return fact
}
