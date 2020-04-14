package main

import "fmt"

//Create a func which returns a func
//assign the returned func to a variable
//call the returned func


func main() {
	a := fooo()
	fmt.Println(a())
}



func fooo() func() int {
	fmt.Println("nothing to do in here")

	return func() int {
		return 2
	}
	//fmt.Println(a(i))
}
