package main




//import "fmt"
//
////A “callback” is when we pass a func into a func as an argument. For this exercise,
////pass a func into a func as an argument
//
//func f2(f func(i int) int) int{
//	return f(5) + 5
//}
//
//func f3(i int) int{
//	return i + 5
//}
//
//func main() {
//	a := f2(f3)
//	fmt.Println(a)
//}

////Create a func which returns a func
////assign the returned func to a variable
////call the returned func
//
//func f1(i int) (func() int){
//	return func() int{
//		return 123
//	}
//}
//func main() {
//	a1 := f1(3)
//	fmt.Println(a1())
//}
//
//
//
//



//Assign a func to a variable, then call that func
//
//func foo(i int) int {
//	return i + 1
//}
//
//func main() {
//	a := foo(5)
//	fmt.Println(a)
//}


////Build and use an anonymous func
//
//func main() {
//	a := func(i int) int{
//		return i
//	}(5)
//	fmt.Println(a)
//}
//


//import (
//	"fmt"
//	"math"
//)
//
////create a type SQUARE
////create a type CIRCLE
////attach a method to each that calculates AREA and returns it
////circle area= π r 2
////square area = L * W
////create a type SHAPE that defines an interface as anything that has the AREA method
////create a func INFO which takes type shape and then prints the area
////create a value of type square
////create a value of type circle
////use func info to print the area of square
////use func info to print the area of circle
//
//type squar struct{
//	l float64
//}
//type curcl struct {
//	r float64
//}
//
//
//func (s squar) area() float64{
//	return s.l * s.l
//}
//
//func (c curcl) area() float64{
//	return 2 * math.Pi * c.r * c.r
//}
//
//type shape interface {
//	area() float64
//}
//
//func info(s shape) float64{
//	return s.area()
//}
//
//func main() {
//	sq := squar{l: 5.1}
//	ci := curcl{r: 4.4}
//
//	ssq := info(sq)
//	cci := info(ci)
//	fmt.Println(ssq)
//	fmt.Println(cci)
//}


////Create a user defined struct with
////the identifier “person”
////the fields:
////	first
////	last
////	age
////attach a method to type person with
////the identifier “speak”
////the method should have the person say their name and age
////create a value of type person
////call the method from the value of type person
//
//
//type person struct {
//	first string
//	last string
//	age int
//}
//
//func (p person) speak() {
//	fmt.Println(p.first, " " ,p.last, " -> ", p.age)
//}
//func main() {
//	p1 := person{age: 5, first:"first", last: "last"}
//	p1.speak()
//}


//
////Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
//
//func main() {
//	defer fmt.Println("last")
//	fmt.Println("first")
//}


//import "fmt"
//
////create a func with the identifier foo that
////	takes in a variadic parameter of type int
////	pass in a value of type []int into your func (unfurl the []int)
////	returns the sum of all values of type int passed in
////create a func with the identifier bar that
////	takes in a parameter of type []int
////	returns the sum of all values of type int passed in
//
//func foo(i ...int) int{
//	sum := 0
//	for _,v := range i{
//		sum = sum + v
//	}
//	return sum
//}
//
//func bar(i []int) int{
//	sum := 0
//	for _,v := range i{
//		sum = sum + v
//	}
//	return sum
//}
//func main() {
//	a := []int{1,2,3}
//	fmt.Println(foo(a...))
//	fmt.Println(bar(a))
//}
//



////Hands on exercise
////create a func with the identifier foo that returns an int
////create a func with the identifier bar that returns an int and a string
////call both funcs
////print out their results
//
//
//func foo() int{
//	return 5
//}
//
//func bar() (int, string){
//	return 1, "blabla"
//}
//
//
//func main(){
//	fmt.Println(foo())
//	fmt.Println(bar())
//}


//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println(recur(3))
//}
//
//func recur(i int) int{
//	if i == 0 {
//		return 1
//	}
//	return i * recur(i - 1)
//
//}

////Closure is when we have “enclosed” the scope of a variable in some code block.
////For this hands-on exercise, create a func which “encloses” the scope of a variable:
//
//func main() {
//	fmt.Println(pri(6))
//	fmt.Println(pri(3))
//}
//
//func pri(i int) int{
//	return i
//}
////A “callback” is when we pass a func into a func as an argument.
////For this exercise,
////pass a func into a func as an argument
//
//
//func callback(f func(i int) int, in int) int{
//	return f(in) + 6
//
//}
//
//func main() {
//	g := func(i int) int{
//		return i + 5
//	}
//
//	fmt.Println(callback(g, 5))
//}
//

//
////Create a func which returns a func
////assign the returned func to a variable
////call the returned func
//
//var a []int
//func foo(xi []int) func(i int) int {
//	//fmt.Println("fist from slice", xi[0])
//	return func(i int) int{
//		return i + 444
//	}
//}
//
//func main() {
//	a = []int{1333, 4, 1, 3}
//	b := foo(a)
//	c := b(a[0])
//	fmt.Println(c)
//}

////Assign a func to a variable, then call that func
//
//func foo(i int) {
//	//fmt.Printf("%T \t %v", i, i)
//	fmt.Println(i)
//}
//
//func main() {
//	a := foo
//	a(5)
//
//
//}
//
//
//

// Build and use an anonymous func

//func main() {
//	func(i int){fmt.Println(i)}(6)
//}
//


//import (
//	"fmt"
//	"math"
//)
//
////create a type SQUARE
////create a type CIRCLE
////attach a method to each that calculates AREA and returns it
////circle area= π r 2
////square area = L * W
////create a type SHAPE that defines an interface as anything that has the AREA method
////create a func INFO which takes type shape and then prints the area
////create a value of type square
////create a value of type circle
////use func info to print the area of square
////use func info to print the area of circle
//
//type square struct {
//	l float64
//}
//type circle struct {
//	r float64
//}
//func (s square) area() float64{
//	return s.l * s.l
//}
//func (s circle) area() float64{
//	return 2 * math.Pi * s.r
//}
//type shape interface  {
//	area() float64
//}
//func info(s shape) float64 {
//	return s.area()
//}
//func main() {
//	s1 := square{
//		l: 5.0,
//	}
//	c1 := circle{
//		r:3.12,
//	}
//	fmt.Println(info(s1))
//	fmt.Println(info(c1))
//
//}
//

// Create a user defined struct with
//the identifier “person”
//the fields:
//first
//last
//age
//attach a method to type person with
//the identifier “speak”
//the method should have the person say their name and age
//create a value of type person
//call the method from the value of type person
//
//type person struct{
//	first string
//	last string
//	age int
//}
//
//
//func (p person) speak() {
//	fmt.Println(p.last, " ,", p.first, " - ", p.age)
//}
//
//func main() {
//	p1 := person{
//		first: "jj",
//		last:  "aa",
//		age:   12,
//	}
//	p1.speak()
//}
//
//


//
//
////Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
//
//func main() {
//	foo()
//}
//
//func foo(){
//	defer fmt.Println("last")
//	fmt.Println("fisrt")
//}


//import "fmt"
//
////create a func with the identifier foo that
////	takes in a variadic parameter of type int
////	pass in a value of type []int into your func (unfurl the []int)
////	returns the sum of all values of type int passed in
////create a func with the identifier bar that
////	takes in a parameter of type []int
////	returns the sum of all values of type int passed in
////
//
//func main() {
//	a := []int{2,3,4}
//	fmt.Println(foo(a...))
//	fmt.Println(baz(a))
//
////	fmt.Println(a)
//}
//
//func foo(x ...int) int {
//	sum := 0
//	for _,v := range x{
//		//v += v
//		sum = sum + v
//	}
//	return sum
//}
//
//func baz(x []int) int {
//	sum := 0
//	for _,v := range x{
//		sum = sum + v
//	}
//	return sum
//}
//



////Hands on exercise
////create a func with the identifier foo that returns an int
////create a func with the identifier bar that returns an int and a string
////call both funcs
////print out their results
//
//func main() {
//	f := foo()
//	b, b1 := bar()
//	fmt.Println(f, b, b1)
//
//}
//
//func foo() int{
//	return 5
//}
//
//func bar() (int, string){
//	return 2, "string"
//
//}