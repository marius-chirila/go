package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	ID        int
	First     string
	FavColors []string
}

type Users struct {
	ID        int
	First     string
	FavColors []string
}

func main() {
	p1 := Person{
		ID:        1234,
		First:     "James",
		FavColors: []string{"Red", "Blue", "Yellow"},
	}
	p2 := Person{
		ID:        56789,
		First:     "Jenny",
		FavColors: []string{"Gray", "Black", "Fuchsia"},
	}

	people := []Person{p1, p2}

	jsons, err := json.Marshal(people)
	if err != nil {
		log.Fatalln("Error", err)
	}
	os.Stdout.Write(jsons)

	var unMarshals []Person
	errors := json.Unmarshal(jsons, &unMarshals)
	if err != nil {
		log.Fatalln("Did not managed to unmarshal", errors)
	}
	fmt.Println("-------------------------------")
	fmt.Println(unMarshals)

}
