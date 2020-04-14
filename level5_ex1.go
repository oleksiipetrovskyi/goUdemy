package main

import "fmt"

//Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
//first name
//last name
//favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.


func main() {
	type person struct {
		first_name string
		last_name string
		favorite_ice []string
	}

	per1 := person{
		first_name:   "person1_first",
		last_name:    "person1_last",
		favorite_ice: []string{"flav1", "flav2"},
	}

	per2 := person{
		first_name:   "person2_first",
		last_name:    "person2_last",
		favorite_ice: []string{"flav11", "flav12"},
	}

	fmt.Println(per1.favorite_ice)
	fmt.Printf("%T\n",per1.favorite_ice )
	fmt.Println(per2)
}
