package main

import "fmt"

//Create a map with a key of TYPE string which is a person’s “last_first” name,
//and a value of TYPE []string which stores their favorite things. Store three records in your map.
//Print out all of the values, along with their index position in the slice.
//
//	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
//`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
//`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

func main() {
	//map[string]int
	x := map[string][]string{}
	fmt.Println(len(x))
	x["bond_james"] = []string{`"Shaken, not stirred"`, `"Martinis"`, `"Women"`}
	x["moneypenny_miss"] = []string{`"James Bond"`, `"Literature"`, `"Computer Science"`}
	x["no_dr"] = []string{`"Being evil"`, `"Ice cream"`, `"Sunsets"`}

	//m := map[string][]string{
	//	`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
	//	`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
	//	`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	//}

	for k, v := range x {
		fmt.Println("whoi = ", k)
		for k1, v1 := range v {
			fmt.Printf("\t %v = %v \n", k1, v1)
		}
	}
}

