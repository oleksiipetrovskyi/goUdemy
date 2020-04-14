// Starting with this code, pull the values off the channel using a select statement

package main

import (
	"fmt"
	"time"

	//"testing/quick"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)
	close(q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive(c <-chan int, q <-chan int) {
	for {
		select {
			case v := <-c: {
				fmt.Println(v)
			}
			case <-q : {
				time.Sleep(4* time.Second)
				fmt.Println("second case")
//				return
			}
		}
	}
fmt.Println("after for")
}
