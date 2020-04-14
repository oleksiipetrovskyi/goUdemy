package main

import "fmt"

// Assign a func to a variable, then call that func

func main() {
	x := foo
	fmt.Printf("%T, %v\n", x, x)
	x()
}

func foo() {
	fmt.Printf("from Func a \t\n")
	//return "from Func a"
}
