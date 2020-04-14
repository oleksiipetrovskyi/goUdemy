package main

import (
	"fmt"
	"runtime"
	"sync"
)

//in addition to the main goroutine, launch two additional goroutines
//each additional goroutine should print something out
//use waitgroups to make sure each goroutine finishes before your program exists

func init() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}



func main() {
	const wgn = 2
	var wg sync.WaitGroup
	wg.Add(wgn)
	go func() {
		fmt.Println("Goroutines 1:", runtime.NumGoroutine())
		fmt.Println("123123")
		wg.Done()
	}()
	go func() {
		fmt.Println("Goroutines 2:", runtime.NumGoroutine())
		fmt.Println("987987987")
		wg.Done()
	}()
	wg.Wait()
}