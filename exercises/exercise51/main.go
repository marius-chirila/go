package main

import (
	"fmt"

	"github.com/marius-chirila/dog"
)

func main() {
	humanYears := 32
	fmt.Println("You'll have this amount of years in dog years:", dog.Years(humanYears))
}
