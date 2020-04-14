package main

import "fmt"

//Closure is when we have “enclosed” the scope of a variable in some code block.
//For this hands-on exercise, create a func which “encloses” the scope of a variable:
//
func main() {
	fmt.Println(foo(5))
	//fmt.Println(i)
	fmt.Println(foo(8))
}


func foo(i int) int{
	return i
}