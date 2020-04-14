package main

import (
	"fmt"
	"math"
)

// create a type SQUARE
//create a type CIRCLE
//attach a method to each that calculates AREA and returns it
//	circle area= π r 2
//	square area = L * W
//create a type SHAPE that defines an interface as anything that has the AREA method
//create a func INFO which takes type shape and then prints the area
//create a value of type square
//create a value of type circle
//use func info to print the area of square
//use func info to print the area of circle


//type square struct{
//	l float64
//}
//type circle struct {
//	r float64
//}
//
//func (s square) area() float64 {
//	return s.l * s.l
//}
//
//func (c circle) area() float64 {
//	return c.r * c.r * math.Pi
//}
//
//type shape interface {
//	area() float64
//}
//
//func info(s shape) {
//	fmt.Println(s.area())
//}
//
//
//
//func main() {
//	ii := circle {
//		r: 2,
//	}
//
//	squ := square {
//		l: 4,
//	}
//	info(ii)
//	info(squ)
//}


type square []float64
type circle float64

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
func  (s square) area() float64 {
	squaref := 1.0
	for _, v := range s{
		squaref = squaref * v
	}
	return squaref

}

func (c circle) area() float64 {
	return float64(math.Pi * c * c)
}


func main() {
	var s square = []float64{2.0 ,4.3}
	info(s)

	var c circle = 5.2
	info(c)
}
