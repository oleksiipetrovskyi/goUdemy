package main

import "fmt"

// Using a COMPOSITE LITERAL:
//create an ARRAY which holds 5 VALUES of TYPE int
//assign VALUES to each index position.
//Range over the array and print the values out.
//Using format printing
//print out the TYPE of the array


func main() {
	var x [5]int
	fmt.Println(x)
	for i := 0; i <= 4; i++ {
		x[i] = i
	}
	fmt.Println(x)

	for _, v := range x {
		fmt.Println(v)
//		fmt.Printf("%T," v)
	}
	fmt.Printf("%T \n", x)
}
