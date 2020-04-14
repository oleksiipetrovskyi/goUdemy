

//// Hands-on exercise #5
////Show the comma ok idiom starting with this code.

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(){ c<- 44}()
	v, ok := <-c
		fmt.Println(v, ok)

	close(c)

	v, ok =<-c
		fmt.Println(v, ok)
}


//
//
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	c := make(chan int)
//	go func() {
//		c <- 33
//	}()
//
//	time.Sleep(3*time.Second)
//	select {
//		case v, ok := <-c : {
//			if ok {
//				fmt.Println("it is OK!  ",ok, v)
//			} else {
//				fmt.Println("it is not OK!  ", ok, v)
//			}
//			a := <-c
//			fmt.Println(a)
//			close(c)
//		}
//
//		case _, ok := <-c: {
//
//			if ok {
//				fmt.Println("it is OK!  ",ok)
//			} else {
//				fmt.Println("it is not OK!  ", ok)
//			}
//		}
//	}
//
//
//	//v, ok =
//	//	fmt.Println(v, ok)
//}
