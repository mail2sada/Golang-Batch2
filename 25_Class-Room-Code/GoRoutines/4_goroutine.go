package main

import (
	"fmt"
	"sync"
	"time"
)

// Generate messages from routine 1
// Display message from routine 2

func main() {
	var messages []string = []string{}
	var mutex sync.Mutex = sync.Mutex{}
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var msgGenerator []string = []string{"Hello", "Hi", "123", "456"}
		for _, val := range msgGenerator {
			mutex.Lock()
			messages = append(messages, val)
			mutex.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for len(messages) == 0 {
			time.Sleep(100 * time.Millisecond)
		}
		for _, val := range messages {
			fmt.Println(val)
		}
	}()
	wg.Wait()

}
