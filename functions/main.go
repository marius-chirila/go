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

func main() {
	foo()
	bar("I'm from bar!")
	fmt.Println(bars("More bars!"))
	fmt.Println(dogYears("Todd", 33))
}
