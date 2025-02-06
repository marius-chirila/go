package main

import "fmt"

func main() {

	cs := make(chan<- int) // send only ints on the channel
	cr := make(<-chan int) // receive only ints on the channel, cannot send
	cb := make(chan int)   // this is a biderectional channel, can receive and send on it

	go func() {
		cs <- 42
	}()
	go func() {
		<-cr
	}()

	go func() {
		cb <- 43
		<-cb
	}()
	fmt.Printf("%T\t send only channel\n", cs)
	fmt.Printf("%T\t receive only channel\n", cr)
	fmt.Printf("%T\t bidirectional channel\n", cb)

}
