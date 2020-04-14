package main

import "fmt"

//Using a COMPOSITE LITERAL:
//create a SLICE of TYPE int
//assign 10 VALUES
//Range over the slice and print the values out.
//Using format printing
//print out the TYPE of the slice


func main() {
	x := []int{ 11, 12, 13, 14, 15,16,17,18,19,20 }

	fmt.Printf("%T\n",x)
	fmt.Println(x)

	for k, v := range x {
		fmt.Printf("%T\n", k)
		fmt.Printf("%T\t %v\n", v, v)
	}

}
