package main

import (
	"fmt"
	"sync"
)

func Display(msg string) {
	defer wg.Done()
	fmt.Println(msg)
}

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	fmt.Println("Demo: Goroutine...")
	wg.Add(2)
	go Display("This is getting printed in a seperate go routine...")

	go Display("This is my second go routine...")

	fmt.Println("Exiting main")
	wg.Wait()
}
