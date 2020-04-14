package main

import "fmt"

//Create a for loop using this syntax
//for { }
//Have it print out the years you have been alive.

func main() {
	a := 1985
	for {
		fmt.Println(a)
		a++
		if a > 2020 {
			break
		}
	}

}
