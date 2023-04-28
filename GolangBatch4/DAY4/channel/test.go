package main

import (
	"fmt"
	"sync"
	"time"
)

func TestWriteChan() {
	defer wg.Done()
	defer close(ch)

	time.Sleep(5 * time.Second)
	fmt.Println("Writing to channel")
	//for i := 0; i < 10; i++ {

	ch <- 10
	fmt.Println("Wrote ", 10)
	//}
}

func TestReadChan() {
	defer wg.Done()

	fmt.Println("Reading from channel")
	val, ok := <-ch
	fmt.Println("Read from channnel", ok, val)
	// for {

	// 	val, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("Read val", val)

	// }
}

var ch chan int
var wg sync.WaitGroup = sync.WaitGroup{}

func main() {

	fmt.Println("Channels")
	ch = make(chan int)
	wg.Add(2)
	go TestWriteChan()
	go TestReadChan()
	wg.Wait()

}
