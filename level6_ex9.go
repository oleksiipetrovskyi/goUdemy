package main

import (
	"fmt"
	"strings"
)

// A “callback” is when we pass a func into a func as an argument. For this exercise,
//pass a func into a func as an argument

func main() {
	xss := []string{"this","is","run","  - >"}

	fmt.Println(callback(toString, xss))
}


func callback(f func(xi []string) string, s []string) string{
	r := f(s)
	ret := fmt.Sprintln(r, "from the callback")
	return ret
}

func toString (xi []string) string {
		return strings.Join(xi, " ")
	}