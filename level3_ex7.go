package main

import "fmt"

//Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

func main() {
	for i:=0; i< 100; i++{
		if i%4 < 2 {
			fmt.Println("reminder of %3  < 2", i)
		} else if i%4 < 3 {
			fmt.Println("reminder of %3  < 3", i)
		} else {
			fmt.Println("asd")
		}

	}
}
