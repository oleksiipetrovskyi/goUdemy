//Starting with this code, pull the values off the channel using a for range loop
package main

import (
	"fmt"
	"time"
)

func main() {

	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)

	}()

	//fmt.Println(<-c)
	return c

}
