package main

import "fmt"

func main() {
	xi := []int{1, 4, 6, 2, 8, 3, 7}

	// xii := make([]int, 7)
	b := false
	for !b {
		b = true
		for i := 0; i < len(xi)-1; i++ {
			if xi[i] > xi[i+1] {
				xi[i], xi[i+1] = xi[i+1], xi[i]
				b = false
			}
		}
	}
	fmt.Println(xi)
}
