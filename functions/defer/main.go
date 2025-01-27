package main

import "fmt"

func main() {
	defer foo() // don't run this function until the end of parent function
	bar()
}

func foo() {
	fmt.Println("I am foo!")
}
func bar() {
	fmt.Println("I am bar!")
}
