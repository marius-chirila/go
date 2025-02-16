package main

import (
	"fmt"
)

type person struct {
	first string
}

func changeField(p person, s string) person {
	p.first = s
	return p
}

func changePointerField(p *person, s string) {
	p.first = s
}

func main() {
	p1 := person{first: "Michael"}
	fmt.Println(p1)
	fmt.Println(changeField(p1, "George"))
	fmt.Println(p1)
	changePointerField(&p1, "James")
	fmt.Println(p1)

}
