package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Demo: channels.")
	wg := sync.WaitGroup{}

	var channel chan string = make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(channel)
		var messageGen []string = []string{"Hi", "Hello", "123", "abc", "Hi", "Hello", "123", "abc"}
		for _, val := range messageGen {
			channel <- val
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range channel {
			fmt.Println(val)
		}
	}()

	wg.Wait()
}
