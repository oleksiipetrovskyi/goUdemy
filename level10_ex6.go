package main

import "fmt"

// write a program that
//puts 100 numbers to a channel
//pull the numbers off the channel and print them

func main()  {
	c := make (chan int)
	go populate(c)
	reader(c)
}


func populate(c chan<- int)  {
	for i := 0; i < 100; i++{
		c<- i
	}
	close(c)
}
func reader(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}