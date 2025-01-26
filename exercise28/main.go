package main

import (
	"fmt"
)

type person struct {
	first           string
	last            string
	iceCreamFlavors []string
}

func printIceCream(x []string) {
	for _, v := range x {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n-------------------------\n")
}

var james = person{
	first:           "James",
	last:            "Bond",
	iceCreamFlavors: []string{"Vanilla", "Chocolate"},
}

var jenny = person{
	first:           "Jenny",
	last:            "Monneywise",
	iceCreamFlavors: []string{"Pistachio", "Salted Caramel"},
}

func main() {
	// fmt.Printf("The first person is %v %v and favorite IceCream flavors are: ", james.first, james.last)
	// printIceCream(james.iceCreamFlavors)
	// fmt.Printf("\nThe second person is %v %v and favorite IceCream flavors are: ", jenny.first, jenny.last)
	// printIceCream(jenny.iceCreamFlavors)
	m := make(map[string]person)
	m[jenny.last] = jenny
	m[james.last] = james
	for _, v := range m {
		fmt.Printf("The person is %v %v and the favorite iceCream Flavors are: ", v.first, v.last)
		printIceCream(v.iceCreamFlavors)
	}
	fmt.Println(m)
}
