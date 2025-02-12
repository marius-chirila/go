package word

import (
	"fmt"
	"testing"

	"github.com/marius-chirila/go/exercise53/quote"
)

func TestCount(t *testing.T) {
	s := "Going down the alley, with my flip-flops on"
	i := Count(s)
	if i != 8 {
		t.Error("Expected", 8, "got", i)
	}
}
func TestUseCount(t *testing.T) {
	s := quote.SunAlso
	m := UseCount(s)
	if m["married"] != 2 {
		t.Error("Expected", 2, "Got", m["married"])
	}

}

func ExampleCount() {
	s := "Going down the alley, with my flip-flops on"
	fmt.Println(Count(s))
	// Output:
	// 8
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < len(quote.SunAlso); i++ {
		Count(quote.SunAlso)
	}

}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < len(quote.SunAlso); i++ {
		UseCount(quote.SunAlso)
	}

}
