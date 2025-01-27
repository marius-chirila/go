package main

import (
	"fmt"
)

type person struct {
	first string
}
type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() { // function associated with a struct -> methods
	fmt.Println("I am ", p.first)
}
func (s secretAgent) speak() { // function associated with a struct -> methods
	fmt.Println("I am a secret agent ", s.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Marcel",
	}
	// p2 := person{
	// 	first: "Aiurea",
	// }

	s1 := secretAgent{
		person: person{
			first: "Jamess",
		},
		ltk: true,
	}
	// p1.speak()
	// p2.speak()
	// s1.speak()
	saySomething(p1)
	saySomething(s1)

}
