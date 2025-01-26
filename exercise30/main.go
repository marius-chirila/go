package main

import (
	"fmt"
)

func main() {

	p1 := struct { // anonymous struct
		first    string
		friends  map[string]int
		favDrink []string
	}{
		first: "Marcel",
		friends: map[string]int{
			"Iuliana": 27,
			"Q":       29,
			"Irina":   30},
		favDrink: []string{"tequila", "rum", "coke"},
	}
	fmt.Println("The name's", p1.first, "and his friends are:")

	for k, v := range p1.friends {
		fmt.Println(k, "and is", v, "old.")
	}
	fmt.Printf("Favorite drinks are:")
	for _, v := range p1.favDrink {
		fmt.Printf(" %v", v)
	}
}
