package main

import (
	"fmt"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47} //slice creation
	for i, value := range xi {
		fmt.Printf("The index is %v and the value is %v\n", i, value)
	}

	m := map[string]int{ // map creation
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("The key is %v \t the value is %v \n", k, v)
	}

}
