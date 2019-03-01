package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println("i: ", i)
	//		wg.Done()
	//	}(i)
	//}
	wg.Wait()
}
