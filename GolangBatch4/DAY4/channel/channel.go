package main

import (
	"fmt"
	"sync"
)

func Procuder(ch chan int, num ...int) {

	go func() {
		defer wg.Done()
		defer close(ch)
		for _, val := range num {
			ch <- val
			fmt.Println("This is from Producer", val)

		}
	}()

}

func Consumer(ch chan int) {
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println("this is from consumer", val)
		}
	}()
}

var wg = sync.WaitGroup{}

func main() {
	var producerChannel chan int = make(chan int)
	fmt.Println("Demo channels")
	wg.Add(2)
	Procuder(producerChannel, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	Consumer(producerChannel)
	wg.Wait()

}
