package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 42,
		"Ethan": 32,
		"Paul":  14,
	}

	if age, ok := m["Paul"]; ok {
		fmt.Println(age)
	} else {
		fmt.Println("Not part of the map")
	}
}
