package main

import (
	"fmt"
)

func main() {
	m := make(map[string][]string)
	v := []string{"shaken", "not stirred", "martinis", "fast cars", "intelligence", "literature", "computer science", "cats", "ice cream", "sunsets"}
	m["bond_james"] = v[:4]
	m["moneypenny_jenny"] = v[4:7]
	m["no_dr"] = v[7:]
	m["feming_ian"] = []string{"steaks", "cigars", "espionage"}
	for k, v := range m {
		fmt.Println(k)
		for i, vv := range v {
			fmt.Printf("Index position is %v and the value is: %v\n", i, vv)
		}
	}
	delete(m, "feming_ian")
	for _, v := range m {
		for _, vv := range v {
			fmt.Printf("New map with values is: %v\n", vv)
		}
	}

}
