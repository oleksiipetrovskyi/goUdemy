package main

import "fmt"

//Building on the code from the previous example
//at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”.
//The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
//eg:
//
//in func main
//this should already be done
//print out the value of the variable “x”
//print out the type of the variable “x”
//assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
//print out the value of the variable “x”
//now do this
//now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
//then use the “=” operator to ASSIGN that value to “y”
//print out the value stored in “y”
//print out the type of “y”
type mytype int

var x mytype
var y int

func main() {

	fmt.Printf("%v\t", x)
	fmt.Printf("%T\t", x)
	x = 42
	fmt.Println("x=", x)
	y = int(x)
	fmt.Printf("%T\t", y)
	fmt.Println("y=", y)
}
