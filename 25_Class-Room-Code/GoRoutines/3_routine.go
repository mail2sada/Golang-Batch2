package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	fmt.Println("Demo: Routines...")
	wg.Add(1)
	go func() { //Closure
		defer wg.Done()
		fmt.Println("This is my go routine...")
	}()
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("This is my second go routine...")
	}()
	wg.Wait()
}
