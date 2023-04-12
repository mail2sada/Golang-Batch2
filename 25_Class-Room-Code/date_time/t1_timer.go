package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	fmt.Println("Demo: Timer using AfterFunc")
	wg.Add(1)
	tmr := time.AfterFunc(5*time.Second, func() {
		defer wg.Done()
		fmt.Println("My timer is expired")
	})

	tmr.Stop()
	wg.Wait()
}
