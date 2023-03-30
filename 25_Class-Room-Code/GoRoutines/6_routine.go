package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(pchan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("In producer")
	time.Sleep(5 * time.Second)
	fmt.Println("After sleep")
	pchan <- 10
	fmt.Println("Wrote 10")
	pchan <- 20
	fmt.Println("Wrote 20")
	pchan <- 30
	fmt.Println("Wrote 30")

}

func consumer(cchan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("In consumer")
	val := <-cchan
	fmt.Println(val)

	time.Sleep(2 * time.Second)

	val = <-cchan
	fmt.Println(val)

	time.Sleep(2 * time.Second)
	val = <-cchan
	fmt.Println(val)
}

func main() {
	fmt.Println("Demo channels...")

	ch := make(chan int)
	defer close(ch)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go producer(ch, &wg)

	go consumer(ch, &wg)
	wg.Wait()

}
