//write a program that
//launches 10 goroutines
//each goroutine adds 10 numbers to a channel
//pull the numbers off the channel and print them

package main

import (
	"fmt"
	"runtime"
)

// decimal count
const De = 10

func main() {
	c := make(chan int)

	for i :=0; i < De; i++ {
		go populate(c, i)
	}
	readerr(c)
}

func populate(c chan<- int, de int) {
	for i :=0; i < 10; i++ {

		curNumber := de * 10 + i
		c<- curNumber

		if curNumber + 1  == De * 10 {
			close(c)
		}
	}
}

func readerr(c <-chan int){
	for v := range c {
		fmt.Println(v)
	}

}