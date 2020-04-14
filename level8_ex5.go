package main

import (
	"fmt"
	"sort"
)

// Starting with this code, sort the []user by
//age
//last
//Also sort each []string “Sayings” for each user
//print everything in a way that is pleasant



type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type ByFirst []user

func (a ByFirst) Len() int {
	return len(a)
}
func (a ByFirst) Less(i, j int) bool {
	return a[i].First < a[j].First
}
func (a ByFirst) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}


type ByAge []user
func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type BySay []string

func (a BySay) Len() int {
	return len(a)
}
func (a BySay) Less(i, j int) bool {
	return a[i] < a[j]
}
func (a BySay) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)




	sort.Sort(ByFirst(users))
	for _, v := range users{
		fmt.Println(v.First)
	}
	sort.Sort(ByAge(users))
	for _, v := range users{
		fmt.Println(v.Age, v.First, v.Last)
		sort.Strings(v.Sayings)
		for _, v1 := range v.Sayings{
			fmt.Printf("\t%v\n", v1)
		}
	}
	//fmt.Println(users)

	// your code goes here

}

