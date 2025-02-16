package main

import (
	"fmt"
)

type customErr struct {
}

func (c *customErr) Error() string {
	return fmt.Sprintln("I'm an error")

}

func main() {
	c := customErr{}
	foo(&c)
}

func foo(e error) {
	fmt.Printf("\nReturning error%v", e)
}
