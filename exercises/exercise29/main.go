package main

import (
	"fmt"
)

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

var audi = vehicle{
	engine: engine{electric: true},
	make:   "Audi",
	model:  "a4",
	doors:  4,
	color:  "navy blue",
}

func main() {
	bmw := vehicle{
		engine: engine{electric: false},
		make:   "BMW",
		model:  "m4i",
		doors:  2,
		color:  "schwarz",
	}

	fmt.Println(audi, bmw)
	fmt.Println("The first car is", audi.make, "and has the following specs: engine electric", audi.electric, audi.model, audi.doors, audi.color)
	fmt.Println("The first car is", bmw.make, "and has the following specs: engine electric", bmw.engine.electric, bmw.model, bmw.doors, bmw.color)

}
