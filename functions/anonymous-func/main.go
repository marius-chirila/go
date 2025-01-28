package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous func, no params")
	}()

	func(s string) {
		fmt.Printf("\nName is %v", s)
	}("James")

	fmt.Println(func(s string) string {
		return fmt.Sprintf("\n%v", s)
	}("Jenny"))

}
