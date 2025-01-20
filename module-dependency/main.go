package main

import (
	"fmt"

	"github.com/marius-chirila/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1, s2)

	fmt.Println(puppy.Bark())
}
