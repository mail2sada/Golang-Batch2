// Fib		-->
// Generator(1,2,3,4,5,6,7,8,9) --> Processor() --> Consumer()
package main

import (
	"fmt"
	"sync"
)

func Producer(numbers ...int) chan int {
	stream := make(chan int)

	go func() {
		defer wg.Done()
		defer close(stream)
		fmt.Println("Producer::In Producer,,,")
		for _, val := range numbers {
			fmt.Println("Producer::", val)
			stream <- val
		}
	}()
	return stream
}

func Processor(in <-chan int) chan []int {

	stream := make(chan []int)

	go func() {
		defer wg.Done()
		defer close(stream)
		fmt.Println("Processor:: In")
		for val := range in {
			fmt.Println("Processor::", val)
			var res []int = []int{}
			for i := 1; i <= val; i++ {
				res = append(res, i)
			}
			fmt.Println("Processor:: Series", res)
			stream <- res
		}
	}()
	return stream
}

func Consumer(in <-chan []int) {
	go func() {
		fmt.Println("Consumer:: in ")
		defer wg.Done()
		for slice := range in {
			fmt.Println("Consumer:: slice")
			for idx, val := range slice {
				fmt.Println(val, idx)
			}
		}
	}()
}

var wg = sync.WaitGroup{}

func main() {

	fmt.Println("Demo pipe")
	wg.Add(3)
	// ch := Producer(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// stream := Processor(ch)

	// Consumer(stream)

	Consumer(Processor(Producer(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)))

	wg.Wait()

}
