package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Fix the race condition you created in the previous exercise by using a mutex
//it makes sense to remove runtime.Gosched()

func main() {

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	lame := 0
	lame2 := 0

	const incr = 50
	var wg sync.WaitGroup
	wg.Add(incr)

	var mut sync.Mutex
	mut.Lock()

	for i := 0 ; i < incr; i ++ {
		//runtime.Gosched()

		fmt.Println("Goroutines from loop:", runtime.NumGoroutine())
		go func() {
			mut.Lock()
			//runtime.Gosched()
			v := lame
			lame2 =lame2 + 1
			fmt.Printf("i : %v \n", i)
			fmt.Println("lame2 from loop:",lame2)
			v++
			lame = v
			//fmt.Println("lame:", lame)
			//fmt.Println("v:", v)
			//fmt.Println("_____")
			mut.Unlock()
			wg.Done()
		}()
	}
	mut.Unlock()
	wg.Wait()
	fmt.Println("lame2: ",lame2)
	fmt.Println("lame last", lame)


}

