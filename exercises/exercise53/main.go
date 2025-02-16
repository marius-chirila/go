package main

import (
	"fmt"

	"github.com/marius-chirila/go/exercises/exercise53/quote"
	"github.com/marius-chirila/go/exercises/exercise53/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
