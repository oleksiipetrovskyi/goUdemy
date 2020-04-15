// Starting with this code, use the sqrt.Error struct as a value of type error.
//If you would like, use these numbers for your
//lat "50.2289 N"
//long "99.4656 W"


package main

import (
	"errors"
	//"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  float64
	long int
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error, struc values: %v, %v; %v", se.lat, se.long, se.err)
}

func main() {
	b := sqrtError{
		lat: -10.23,
		long: 3,
		err: errors.New(`Error in b struct:`),

	}
	_, err := sqrt(b)
	if err != nil {
		log.Println(err)
	}

}

func (e *sqrtError) Modify(message error) {
	e.err =  message
}


func sqrt(f sqrtError ) (float64, error) {
	if f.lat < 0 {

		e := fmt.Errorf(`%v only positie values allowed for sqrt func, got: lat = "%v"`, f.err, f.lat)
		f.Modify(e)

		return float64(0), f
		//
	}
	return 42, nil
}
