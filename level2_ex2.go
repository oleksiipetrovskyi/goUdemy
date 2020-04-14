package main

import "fmt"

// Using the following operators, write expressions and assign their values to variables:
//==
//<=
//>=
//!=
//<
//>
//Now print each of the variables.

func main() {

	a := 33 == 33
	b := 33 <= 33
	c := 33 >= 33
	d := 33 != 33
	e := 33 < 33
	f := 33 > 33

	fmt.Printf("%v\t, %v\t, %v\t, %v\t, %v\t, %v\n", a, b, c, d, e, f)

}
