package main

import (
	Dog "firstTry/level12/dog"
	"fmt"
	"log"
)

func main() {
	var human int
	fmt.Println("Human years (type int): ")

	_, err := fmt.Scanln( &human )
	if err != nil{
		log.Fatalf("Error in user input: %v", err)
	}
	fmt.Printf("Dog years: %v \n", Dog.Years(human))
}
