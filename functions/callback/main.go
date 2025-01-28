package main

import "fmt"

func main() {
	x := doMath(42, 32, add)
	fmt.Println("Numbers added:", x)
	y := doMath(42, 32, subtract)
	fmt.Println("Numbers subtracted:", y)

}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func doMath(a, b int, f func(int, int) int) int { // callback,
	return f(a, b)
}
