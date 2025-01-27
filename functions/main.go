package main

// func (r receiver) identifier(p parameters(s)) (return (s)) { <code> }

import (
	"fmt"
)

func foo() { //no params, no returns
	fmt.Println("Foo")
}

func bar(s string) { // one param, no return
	fmt.Println(s)
}

func bars(s string) string { // one param, one return
	return s
}

func dogYears(name string, ageDogYears int) (string, int) {
	ageDogYears *= 7
	return fmt.Sprintf("My name is: %v and I would be in dog years:", name), ageDogYears
}

func variadicFunc(x ...int) int { //variadic func with variadic parameters "..."
	fmt.Printf("%v, %T", x, x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	xi := []int{42, 44, 46}
	foo()
	bar("I'm from bar!")
	fmt.Println(bars("More bars!"))
	fmt.Println(dogYears("Todd", 33))
	sum := variadicFunc(42, 33, 54)
	fmt.Println("Sum is: ", sum)
	sum2 := variadicFunc(xi...)
	fmt.Println("New sum is: ", sum2)
}
