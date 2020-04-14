package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Using goroutines, create an incrementer program
//have a variable to hold the incrementer value
//launch a bunch of goroutines
//each goroutine should
//read the incrementer value
//store it in a new variable
//yield the processor with runtime.Gosched()
//increment the new variable
//write the value in the new variable back to the incrementer variable
//use waitgroups to wait for all of your goroutines to finish
//the above will create a race condition.
//Prove that it is a race condition by using the -race flag
//if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej

func main() {

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	lame := 0

	const incr = 50
	var wg sync.WaitGroup
	wg.Add(incr)


	for i := 0 ; i < incr; i ++ {
		go func() {
			runtime.Gosched()
			v := lame
			fmt.Printf("incr: %v \n", i)
			v++
			lame = v
			fmt.Println("lame:", lame)
			fmt.Println("v:", v)
			fmt.Println("_____")

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("lame last", lame)


}

