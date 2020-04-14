package main

import "fmt"

//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
//
//"James", "Bond", "Shaken, not stirred"
//"Miss", "Moneypenny", "Helloooooo, James."
//
//Range over the records, then range over the data in each record.



func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	c := [][]string{a, b}
	//fmt.Println(c)

	for k, v := range c {
		fmt.Printf("first key = %v\t\n", k)
		fmt.Printf("%T\n", v)
		fmt.Printf("%T\n", k)

		for k1, v1 := range v{
			fmt.Printf("\tsecond key = %v\t %v\n", k1, v1)
			fmt.Printf("%T", v1)
		}
	}

}
