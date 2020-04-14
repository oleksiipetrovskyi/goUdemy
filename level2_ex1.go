package main

import "fmt"

// Write a program that prints a number in decimal, binary, and hex

func main() {
	num := 32

	fmt.Printf("%#x\n", num)
	fmt.Printf("%d\n", num)
	fmt.Printf("%b\n", num)
}
