package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Println(a, b, c, d)
	fmt.Printf("\n%T\t%T\t%T\t%T\n", a, b, c, d)
	fmt.Println(*a, *b, *c, *d)

}
