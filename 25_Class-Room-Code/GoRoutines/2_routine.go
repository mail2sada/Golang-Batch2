package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

var mutex = sync.Mutex{}
var counter int = 0

func main() {
	wg.Add(2)

	go myRoutine("One:")
	go myRoutine("Two:")
	wg.Wait()
	fmt.Println("Value of counter...", counter)
}

func myRoutine(str string) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}
