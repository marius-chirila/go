package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct { // embedded struct
	p   person
	ltk bool
}

func main() {

	p2 := struct { //anonymous struct, useful for 1 time use, not recurring like creating a differnt type
		first string
		last  string
		age   int
	}{
		first: "Jenny",
		last:  "Monneywise",
		age:   33,
	}
	sa1 := secretAgent{
		p: person{
			first: "James",
			last:  "Bond",
			age:   23,
		},
		ltk: true,
	}
	fmt.Println(sa1.p.last, sa1.p.first, sa1.p.age, sa1.ltk)
	fmt.Printf("The name's: %v %v, and I'm %v yo. Created by %T\n", p2.first, p2.last, p2.age, p2)
}
