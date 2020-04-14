package main

import "fmt"

// Create a user defined struct with
//	- the identifier “person”
//	  the fields:
//		first
//		last
//		age
//- attach a method to type person with
//		the identifier “speak”
//		the method should have the person say their name and age
//- create a value of type person
//- call the method from the value of type person

type person struct {
	first string
	last string
	age int
}

func (p person) speak()  {
	fmt.Println("name =", p.first, " ", p.last, ", age is - ", p.age)
}


func main() {
	p1 := person{
		first: "james",
		last:  "bond",
		age:   55,
	}
	//fmt.Printf("%T\n", p1.speak)
	p1.speak()

}
