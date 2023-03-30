package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Go routine")

	go fmt.Println("My go routine...")

	go fmt.Println("This is my tes routine...")

	time.Sleep(100 * time.Millisecond)
}
