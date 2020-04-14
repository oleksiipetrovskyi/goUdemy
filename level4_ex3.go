package main

import "fmt"

//Using the code from the previous example, use SLICING to create the following new slices which are then printed:
//[42 43 44 45 46]
//[47 48 49 50 51]
//[44 45 46 47 48]
//[43 44 45 46 47]

func main() {
//	x := []int{ 11, 12, 13, 14, 15,16,17,18,19,20 }
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}
