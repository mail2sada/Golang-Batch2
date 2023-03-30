package main

import (
	"fmt"
	"sync"
	"time"
)

var cond sync.Cond //= sync.NewCond()

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var mutex sync.Mutex = sync.Mutex{}
	cond = *sync.NewCond(&mutex)
	fmt.Println("Demo Conditional Variables")
	var condition = false

	go func() {
		defer wg.Done()
		fmt.Println("In go routine")
		cond.L.Lock()
		for condition == false {
			fmt.Println("Waiting for condition")
			cond.Wait()
		}
		cond.L.Unlock()
		fmt.Println("Exiting routine..")
	}()

	go func() {
		defer wg.Done()

		fmt.Println("In go routine")
		cond.L.Lock()
		for condition == false {
			fmt.Println("Waiting for condition")
			cond.Wait()
		}
		cond.L.Unlock()
		fmt.Println("Exiting routine..")
	}()
	fmt.Println("Sleeping for 10")
	time.Sleep(10 * time.Second)
	fmt.Println("Condition meet")
	cond.L.Lock()
	condition = true
	fmt.Println("Signaling...")
	cond.Broadcast()
	cond.L.Unlock()
	wg.Wait()

}
