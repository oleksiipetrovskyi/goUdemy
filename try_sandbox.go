package main

import "fmt"

func main() {

	s := struct {
		doors int
		color string
	}{
		doors: 2,
		color: "blue",
	}
	fmt.Println(s)


}

//Create and use an anonymous struct.
//
//import "fmt"
//
////Create a new type: vehicle.
////The underlying type is a struct.
////The fields:
////doors
////color
////Create two new types: truck & sedan.
////The underlying type of each of these new types is a struct.
////Embed the “vehicle” type in both truck & sedan.
////Give truck the field “fourWheel” which will be set to bool.
////Give sedan the field “luxury” which will be set to bool. solution
////Using the vehicle, truck, and sedan structs:
////using a composite literal, create a value of type truck and assign values to the fields;
////using a composite literal, create a value of type sedan and assign values to the fields.
////Print out each of these values.
////Print out a single field from each of these values.
//
//func main() {
//	type vehicle struct {
//		doors int
//		color string
//	}
//	type truck struct {
//		vehicle
//		fourWheel bool
//	}
//	type sedan struct {
//		vehicle
//		luxury bool
//	}
//
//	tr := truck{
//		vehicle:   vehicle{
//			doors: 2,
//			color: "black",
//		},
//		fourWheel: true,
//	}
//	sed := sedan{
//		vehicle: vehicle{
//			doors: 4,
//			color: "while",
//		},
//		luxury: false,
//	}
//	fmt.Println(tr, sed)
//	fmt.Println(tr.doors, sed.doors)
//
//}
//





//import "fmt"
//
////Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
////first name
////last name
////favorite ice cream flavors
////Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
//
//
////Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
////Access each value in the map. Print out the values, ranging over the slice.
//
//
//func main() {
//	type persone struct {
//		first_name string
//		last_name string
//		favorite []string
//	}
//
//	p1 := persone{
//		first_name: "james",
//		last_name:  "bond",
//		favorite:   []string{"fav1","fav2","fav3"},
//	}
//
//	p2:= persone{
//		first_name: "ilon",
//		last_name:  "musk",
//		favorite:   []string{"fav11", "fav12", "fav13"},
//	}
//
//	for k,v := range p1.favorite {
//		fmt.Printf("\t %v %v\n", k, v)
//	}
//	fmt.Println(p2)
//
//	m := map[string]persone{
//		p1.last_name: p1,
//		p2.last_name: p2,
//	}
//	fmt.Println(m)
//
//	for k,v := range m {
//		fmt.Printf("who = %v %v\n", k, v.first_name)
//		for _, v1 := range v.favorite {
//			fmt.Println(v1)
//		}
//	}
//}
//
//


//
//import "fmt"
//
//// Create a map with a key of TYPE string which is a person’s “last_first” name,
////and a value of TYPE []string which stores their favorite things. Store three records in your map.
////Print out all of the values, along with their index position in the slice.
////
//
//
////Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
//
////Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
//
////
////
////	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
////`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
////`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
//func main() {
//	//m := map[string][]string{
//	//	`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
//	//	`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
//	//	`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
//	//}
//
//	 x := map[string][]string{}
//	 x["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
//	 x["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
//	 x["no_dr"]  = []string{`Being evil`, `Ice cream`, `Sunsets`}
//	 x["fleming_ian"] = []string{`steaks`, `cigars`, `espionage`}
//
//	fmt.Println(x)
//	for k, v := range x{
//		fmt.Printf("who:%v\n", k)
//		for k1, v1 := range v{
//			fmt.Printf("\t %v - %v\n", k1, v1)
//		}
//	}
//	delete(x, "no_dr")
//	for k, v := range x{
//		fmt.Printf("who:%v\n", k)
//		for k1, v1 := range v{
//			fmt.Printf("\t %v - %v\n", k1, v1)
//		}
//	}
//
//}
//
//
//

//import "fmt"
////Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
////
////"James", "Bond", "Shaken, not stirred"
////"Miss", "Moneypenny", "Helloooooo, James."
////
////Range over the records, then range over the data in each record.
////
//
//func main() {
//	//x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."} }
//	x1 := []string{"James", "Bond", "Shaken, not stirred"}
//	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
//
//	x := [][]string{x1, x2}
//	fmt.Println(x)
//	for k, v := range x{
//		fmt.Println(k, v)
//		for k1, v1 := range v {
//			fmt.Printf("\t%v %v\n", k1 +1 , v1)
//		}
//	}
//}




//import "fmt"
//
////Create a slice to store the names of all of the states in the United States of America.
////What is the length of your slice? What is the capacity?
////Print out all of the values, along with their index position in the slice, without using the range clause.
////Here is a list of the states:
////` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
//
//func main() {
//	x := make([]string, 50, 50)
//	fmt.Println(cap(x), len(x))
//	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, " Wyoming",)
//	fmt.Println(x)
//	fmt.Println(cap(x), len(x))
//	for i := 0; i < len(x) ; i++  {
//		fmt.Println(i, x[i])
//	}
//
//}

//
//import "fmt"
//
////To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
////start with this slice
////x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
////use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
////[42, 43, 44, 48, 49, 50, 51]
//
//func main() {// 0   1   2   3   4   5   6   7  8   9
//	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//	y := append(x[:3], x[6:]...)
//	fmt.Println(y)
//}
//


//import "fmt"
//
////Follow these steps:
////start with this slice
////x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
////append to that slice this value
////52
////print out the slice
////in ONE STATEMENT append to that slice these values
////53
////54
////55
////print out the slice
////append to the slice this slice
////y := []int{56, 57, 58, 59, 60}
////print out the slice
//
//func main() {
//	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//	x = append(x, 52)
//	fmt.Println(x)
//	x = append(x, 53,54,55)
//	fmt.Println(x)
//	y := []int{56, 57, 58, 59, 60}
//	fmt.Println(append(x, y...))
//}



////Using the code from the previous example, use SLICING to create the following new slices which are then printed:
////[42 43 44 45 46]
////[47 48 49 50 51]
////[44 45 46 47 48]
////[43 44 45 46 47]
//
//func main() {
//	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//	y := x[:5]
//	fmt.Println(y)
//	z := x[5:]
//	fmt.Println(z)
//	w := x[2:7]
//	fmt.Println(w)
//	fmt.Println(x[1:6])
//
//}





//import "fmt"
//
////Using a COMPOSITE LITERAL:
////create a SLICE of TYPE int
////assign 10 VALUES
////Range over the slice and print the values out.
////Using format printing
////print out the TYPE of the slice
//
//func main() {
//	x := []int{}
//	for i := 0; i < 16; i++ {
//		x = append(x, i + 1)
//	}
//	fmt.Println(x)
//	for  k, v := range x {
//		fmt.Printf("index %v = %v\n", k, v)
//	}
//	fmt.Printf("%T\n", x)
//	fmt.Println(cap(x))
//	fmt.Println(len(x))
//}





////Using a COMPOSITE LITERAL:
////create an ARRAY which holds 5 VALUES of TYPE int
////assign VALUES to each index position.
////Range over the array and print the values out.
////Using format printing
////print out the TYPE of the array
//
//func main() {
//
//	var x [5]int
//	fmt.Println(x)
//	for i := 0; i < len(x); i++ {
//		x[i] = i + 1
//	}
//	fmt.Println(x)
//	for i := range x {
//		fmt.Println(x[i])
//	}
//	fmt.Printf("%T\n", x)
//}



//Level4 ^^



//
//import "fmt"
////Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
//
//var favSport string
//
//
//func main() {
//	favSport = "foot"
//	switch favSport {
//	case "foot":
//		fmt.Println("foot123")
//	}
//
//
//}
//
////Create a program that uses a switch statement with no switch expression specified.
//
////func main() {
////	a := 3221
////	switch  {
////	case a == 123:
////		fmt.Println("123")
////	case a == 321:
////		fmt.Println("321")
////	default:
////		fmt.Println("defalt")
////	}
////}
//
//// Create a program that shows the “if statement” in action.
////Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
////func main() {
////	a := 11
////	if a < 10 {
////		fmt.Println("a<10   ...", a)
////	}else if a > 10 {
////		fmt.Println("a>10   ...", a)
////	} else {
////		fmt.Println("a=10   ...", a)
////
////	}
////
////}
//
////Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
////func main() {
////	for i := 10; i <= 100; i++ {
////		fmt.Println( i%4 )
////	}
////}
//
//
////Create a for loop using this syntax
////for { }
////Have it print out the years you have been alive.
////const(
////	a = 1985 + iota
////	b
////
////)
////
////func main() {
////	for {
////		fmt.Println(a)
////		fmt.Println(b)
////
////		if a > 2020 {
////			break
////		}
////
////	}
////}
//
//
//
//
//
////Create a for loop using this syntax
////for condition { }
////Have it print out the years you have been alive.
////func main() {
////	a := 1985
////	for a <= 2020 {
////		fmt.Println(a)
////		a++
////	}
////}
//
//
//
//
////Print every rune code point of the uppercase alphabet three times. Your output should look like this:
////65
////	U+0041 'A'
////	U+0041 'A'
////U+0041 'A'
////66
////	U+0042 'B'
////	U+0042 'B'
////	U+0042 'B'
//// … through the rest of the alphabet characters
////
////
////func main() {
////	for i := 65; i <= 90; i++ {
////		fmt.Println(i)
////		for j := 0; j < 3; j++ {
////			fmt.Printf("\t%#U\n", i)
////		}
////	}
////}
