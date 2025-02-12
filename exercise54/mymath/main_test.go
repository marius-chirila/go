package mymath

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer float64
}

func TestCenteredAvg(t *testing.T) {
	tests := []test{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{5, 23, 44, 15, 2}, 14.333333333333334},
		{[]int{55, 33, 11, 2, 44}, 29.333333333333332},
	}
	for _, v := range tests {
		if CenteredAvg(v.data) != v.answer {
			t.Error("Got", CenteredAvg(v.data), "expected", v.answer)
		}
	}
}

func ExampleCenteredAvg() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println(CenteredAvg(data))
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{55, 33, 11, 2, 44})
	}

}
