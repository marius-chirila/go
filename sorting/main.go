package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {

	p1 := person{"James", 32}
	p2 := person{"Jenny", 27}
	p4 := person{"M", 56}
	p3 := person{"Q", 64}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
