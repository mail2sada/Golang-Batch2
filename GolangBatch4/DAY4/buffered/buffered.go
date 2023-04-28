package main

import (
	"fmt"
	"sync"
	"time"
)

func Producer(wg *sync.WaitGroup, ch chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		var stream = []string{"Hello", "Welcome", "To", "Golang", "training", "Mavenir"}
		for _, val := range stream {

			ch <- val
			fmt.Println("Wrote ", val)
		}
	}()
}

func Consumer(wg *sync.WaitGroup, ch chan string) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		time.Sleep(5 * time.Second)
		for val := range ch {
			fmt.Println("data received on channel is ", val)
			time.Sleep(1 * time.Second)
		}
	}()
}

func main() {

	fmt.Println("Buffered channel")
	wg := sync.WaitGroup{}
	ch := make(chan string, 3)
	//defer close(ch)
	Producer(&wg, ch)
	Consumer(&wg, ch)
	wg.Wait()
}
