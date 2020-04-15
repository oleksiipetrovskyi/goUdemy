// Create a struct “customErr” which implements the builtin.error interface.
//Create a func “foo” that has a value of type error as a parameter.
//Create a value of type “customErr” and pass it into “foo”.
//If you need a hint, here is one.

package main

import (
	"fmt"
	"log"
)

type persone struct {
	first string
	error string
}

func (re persone) Error() string {
	return fmt.Sprintf("error interface message: %v ", re.error)
}

func foo(e error) {
	if e != nil {
		log.Println("error: ", e)
	}
}


func main() {
	example := persone{
		error: "error from persone",
		first: "Name",
	}
	foo(example)
}