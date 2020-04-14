package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//Fix the race condition you created in exercise #4 by using package atomic


func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var lame int64
	var i int64
	const incr = 50
	var wg sync.WaitGroup
	wg.Add(incr)

	for i = 0 ; i < incr; atomic.AddInt64(&i, 1){
		fmt.Println("Goroutines from loop:", runtime.NumGoroutine())
		go func() {
			runtime.Gosched()
			atomic.AddInt64(&lame, 1)
			fmt.Printf("i : %v \n", atomic.LoadInt64(&i))
			fmt.Println("lame from loop:",atomic.LoadInt64(&lame))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("lame last", lame)


}