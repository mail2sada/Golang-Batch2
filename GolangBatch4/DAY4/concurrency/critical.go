package main

import (
	"fmt"
	"sync"
)

var count = 0

func increment() {
	defer wg.Done()
	for idx := 0; idx < 10000; idx++ {
		mutex.Lock()
		count++
		mutex.Unlock()
	}
}

var wg = sync.WaitGroup{}
var mutex = sync.Mutex{}

func main() {

	for i := 0; i < 2; i++ {
		wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	for idx := 0; idx < 10000; idx++ {
		// 		mutex.Lock()
		// 		count++
		// 		mutex.Unlock()
		// 	}
		// }()
		go increment()
	}
	wg.Wait()
	fmt.Println(count)
}
