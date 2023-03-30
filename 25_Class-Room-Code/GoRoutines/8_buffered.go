package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
generator1
			Processeor output
generator2
*/

func generator(out chan<- int, numbers ...int) {
	log.Println("In Generator")
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Generator Routine...")
		for _, val := range numbers {
			out <- val
		}
	}()

}

func processor(in <-chan int) chan int {
	//make a square
	processChan := make(chan int)
	wg.Add(1)
	go func() {
		defer close(processChan)
		defer wg.Done()

		log.Println("In processor routine")
		for val := range in {
			processChan <- val * val
		}
		fmt.Println("Processor:Exit")
	}()
	return processChan
}

func consumer(in <-chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Inconsumer routine")
		for val := range in {
			fmt.Println(val)
		}

	}()
}

var wg = sync.WaitGroup{}

func executor() {

}

func main() {

	fmt.Println("Demo: Buffered channel")
	var bufferedChan = make(chan int, 10)
	defer func() {
		fmt.Println("Closing bufferedChan")
		//close(bufferedChan)
	}()

	generator(bufferedChan, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19)
	generator(bufferedChan, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20)
	consumer(processor(bufferedChan))

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("In async after...")
		close(bufferedChan)
	})
	wg.Wait()
}
