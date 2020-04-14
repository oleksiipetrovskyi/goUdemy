package main

import "fmt"

//本語
//func main() {
//	chi := "日本語"
////	chi := "23 32 333"
//	//for i :=0; i < len(chi); i++ {
//	//	fmt.Printf("%x ", chi[i])
//	//}
//	for i, v := range chi {
//		fmt.Println(i)
//		fmt.Printf("%#x\n", v)
//}
//
//}
type weekday = int
const (
	sunday weekday = iota + 1
	_
    monday
	//a = iota
	//b = 1 << 2


)


func main() {
	//fmt.Printf("%T\n%T", a, b)
	//fmt.Println("a=", a)
	//fmt.Println("b=", b)
	//
	fmt.Printf("sun=%v, mon=%v", sunday, monday)

}