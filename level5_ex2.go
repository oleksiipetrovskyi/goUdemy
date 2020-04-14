package main

import "fmt"

//Take the code from the previous exercise,
//then store the values of type person in a map with the key of last name.
//Access each value in the map. Print out the values, ranging over the slice.

func main() {
	type person struct {
		first_name   string
		last_name    string
		favorite_ice []string
	}

	per1 := person{
		first_name:   "person1_first",
		last_name:    "person1_last",
		favorite_ice: []string{
				"flav1",
				"flav2",
		},
	}

	per2 := person{
		first_name:   "person2_first",
		last_name:    "person2_last",
		favorite_ice: []string{
				"flav11",
				"flav12",
		},
	}
	x := map[string]person{
		per1.last_name: per1,
		per2.last_name: per2,
	}

	fmt.Println(x)
	for _, v := range x {
		fmt.Printf("last = %v\t\n", v)
		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for _, v1 := range v.favorite_ice {
			fmt.Printf("\t%v\n", v1)
		}
	}

}
