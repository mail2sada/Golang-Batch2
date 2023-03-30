package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("Allocating buffer...")
		return new(bytes.Buffer)
	},
}

func log(message string) {
	b := bufPool.Get().(*bytes.Buffer)

	b.Reset()

	b.WriteString(time.Now().GoString())
	b.WriteString(message)

	fmt.Println(string(b.Bytes()))

	bufPool.Put(b)
}

func main() {
	fmt.Println("Demo: Sync Pool")
	log("Hello...")
	log("How are you")

}
