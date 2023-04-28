package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Demo: Goroutine closure...")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Running in closure")
	}()
	fmt.Println("Waiting for go routine to exit")
	wg.Wait()
	fmt.Println("Exiting main")
}
