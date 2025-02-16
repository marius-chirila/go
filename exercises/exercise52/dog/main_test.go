package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(4)
	if x != 28 {
		t.Error("We got", x, "and we expected", 4*7)

	}
}

func ExampleYearsTwo() {
	x := YearsTwo(4)
	fmt.Println(x)
	//Output:
	//28
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(i)
	}
}
