package main

import "fmt"

// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
const (
	current int = 2020
	year = current + iota
	next
	more
)
func main() {
	fmt.Println("current=", current)
	fmt.Println("year=", year)
	fmt.Println("next=", next)
	fmt.Println("more=", more)
}
