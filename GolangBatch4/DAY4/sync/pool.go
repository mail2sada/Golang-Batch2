package main

import (
	"fmt"
	"sync"
)

var pool = sync.Pool{New: func() any {
	fmt.Println("Allocating buffer")
	buffer := make([]byte, 10*1024)
	return buffer
}}

func log(msg string) {
	// appending time stamp
	// appending file name
	buffer := pool.Get()

	// use the buffer
	fmt.Println("", msg)

	pool.Put(buffer)
}

func main() {
	fmt.Println("Demo sync pool")
	log("1")
	log("2")
	log("3")
}
