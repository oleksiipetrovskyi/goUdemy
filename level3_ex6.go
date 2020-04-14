package main

import (
	"fmt"
)

//Create a program that shows the “if statement” in action.
func main() {
	for i:=0; i< 100; i++{
		if i%3 < 2 {
			fmt.Println("reminder of %3  < 2", i)
		} else {
			fmt.Println("asd")
		}

	}
}
