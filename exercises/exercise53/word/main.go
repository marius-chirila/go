package word

import (
	"strings"
)

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words found in the string given via the length of the slice
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
