package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var slice = []string{}

func Producer(cn *sync.Cond) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		slice = append(slice, "A"+strconv.Itoa(i))
	}
	cn.L.Lock()
	//cn.Signal()
	cn.Broadcast()
	cn.L.Unlock()
}

func Consumer(cn *sync.Cond) {
	defer wg.Done()
	for len(slice) == 0 {

		cn.L.Lock()
		fmt.Println("Waiting on lock")
		cn.Wait()
		fmt.Println("got signal")
		cn.L.Unlock()
	}
	for _, val := range slice {
		fmt.Println(val)
	}
}

var wg = sync.WaitGroup{}

func main() {
	var mutex = sync.Mutex{}
	var cond = sync.NewCond(&mutex)
	fmt.Println("Conditional Variable")
	wg.Add(2)
	go Consumer(cond)
	//go Consumer(cond)
	time.Sleep(3 * time.Second)
	go Producer(cond)
	wg.Wait()
}
