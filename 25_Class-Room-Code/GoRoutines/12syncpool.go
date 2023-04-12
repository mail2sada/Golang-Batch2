package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

func log(msg string) {
	var b = pool.Get().(bytes.Buffer)


	b.Reset()

	b.WriteString(time.Now().String())
	b.WriteString(msg)

	fmt.Println(string(b.Bytes()))
	pool.Put(b)
}

var pool = sync.Pool{
	New: func() any {
		fmt.Println("Allocating buffer")
		var b bytes.Buffer
		return b
	},
}

func main() {
	log("Inside main...")
	log("Executing main")
	log("Exiting main")
}
