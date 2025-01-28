package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))

}

func main() {
	p1 := person{
		first: "James",
	}
	f, err := os.Create("output.txt") // create the file
	if err != nil {                   // catch err
		log.Fatalf("%s, error", err)
	}
	defer f.Close() // defer close file

	var b bytes.Buffer

	p1.writeOut(f)
	p1.writeOut(&b)
	fmt.Println(b.String())
}
