package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	gs := 10
	c := make(chan int)
	wg.Add(gs)

	func() {
		for i := 0; i < gs; i++ {
			go func() {
				fmt.Println("Number routines:", runtime.NumGoroutine())
				c <- i
				// close(c)
				wg.Done()
			}()

		}

	}()

	for k := 0; k < 10; k++ {
		fmt.Println(<-c)
	}
	// for v := range c {
	// 	fmt.Println(v)
	// }

	wg.Wait()
	close(c)
}
