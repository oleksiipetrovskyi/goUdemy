package main

import "fmt"

//Create a program that uses a switch statement with no switch expression specified.


func main() {
	a := 4
	switch {
	case a == 4:
		fmt.Println("case = 4")
	}
}
