package main

import "fmt"

// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

func main() {
	favSport := "badminton"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "badminton":
		fmt.Println("go to the court!")
	
	}
}
