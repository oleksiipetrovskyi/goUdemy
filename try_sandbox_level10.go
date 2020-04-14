// Show the comma ok idiom starting with this code.

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(){
		c<- 33
	}()

	v, ok := <-c
		fmt.Println(v, ok)

	close(c)

	v, ok = <-c
		fmt.Println(v, ok)
}




//// Starting with this code, pull the values off the channel using a select statement
//
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	q := make(chan int)
//	c := gen(q)
//
//	receive(c, q)
//
//	fmt.Println("about to exit")
//}
//
//func receive(c, q <- chan int) {
//	for {
//		select {
//		case v := <-c:
//			{
//				fmt.Println(v)
//			}
//		case <-q:
//			{
//				return
//			}
//		}
//	}
//}
//
//func gen(q chan<- int) <-chan int {
//	c := make(chan int)
//
//	go func(){
//		for i := 0; i < 3; i++ {
//			c <- i
//		}
//		q <- 1
//		//close(q)
//		close(c)
//	}()
//
//	return c
//}
//
//
//
//
//
////// Starting with this code, pull the values off the channel using a for range loop
////
////
////package main
////
////import (
////	"fmt"
////)
////
////func main() {
////	c := gen()
////	receive(c)
////
////	fmt.Println("about to exit")
////}
////
////func gen() <-chan int {
////	c := make(chan int)
////
////	go func(){
////		for i := 0; i < 10; i++ {
////			c <- i
////		}
////		close(c)
////	}()
////	return c
////}
////
////func receive(c <- chan int){
////	for v := range c{
////	fmt.Println(v)
////	}
////}
////
//////package main
//////
//////import (
//////	"fmt"
//////)
//////
//////func main() {
//////	cs := make(chan int)
//////
//////	go func() {
//////		cs <- 42
//////	}()
//////	fmt.Println(<-cs)
//////
//////	fmt.Printf("------\n")
//////	fmt.Printf("cs\t%T\n", cs)
//////}
//////
//////
////////package main
////////
////////import (
////////	"fmt"
////////)
////////
////////func main() {
////////
////////	c := make(chan int,1 )
////////
////////	//go func(){
////////		c <- 42
////////	//}()
////////
////////	fmt.Println(<-c)
////////}
