package main

import "fmt"

// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.


func main() {
	defer fmt.Println("defer")
	fmt.Println("first")

}

