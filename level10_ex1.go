package main

import (
"fmt"
	//"sync"
)


func main() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}

