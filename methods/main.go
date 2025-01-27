package main

import (
	"fmt"
)

type person struct {
	first string
}

func (p person) speak() { // function associated with a struct -> methods
	fmt.Println("I am ", p.first)
}

func main() {
	p1 := person{
		first: "Marcel",
	}
	p2 := person{
		first: "Aiurea",
	}

	p1.speak()
	p2.speak()
}
