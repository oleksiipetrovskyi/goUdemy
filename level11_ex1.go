// Start with this code. Instead of using the blank identifier, make sure the code is checking and handling the error.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func check(e error) {
	if e == nil {
		log.Println("error: ", e)
	}
}

func main() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)

	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	check(err)

	fmt.Println(string(bs))

}
