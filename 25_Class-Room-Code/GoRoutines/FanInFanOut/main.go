package main

import (
	"fmt"
	"sync"
)

func Generator(number ...int) chan int {

	out := make(chan int)

	go func() {
		defer waitGroup.Done()

		defer close(out)
		fmt.Println("Generator")
		for _, val := range number {
			out <- val
		}
	}()
	return out
}

func Merge(chanList ...chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	// We have received a list of channels...
	output := func(in chan int) {
		defer wg.Done()
		for val := range in {
			out <- val
		}
	}
	wg.Add(len(chanList))
	for _, ch := range chanList {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Process(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer waitGroup.Done()

		defer close(out)
		for val := range in {
			out <- val * val
		}
	}()
	return out
}

func Consumer(in <-chan int) {
	go func() {
		defer waitGroup.Done()
		for val := range in {
			fmt.Println(val)
		}
	}()
}

var waitGroup = sync.WaitGroup{}

func main() {
	waitGroup.Add(4)
	fmt.Println("Demo FanInFanOut")
	ch1 := Generator(1, 2, 3, 4, 5)
	ch2 := Generator(6, 7, 8, 9, 10)
	merge := Merge(ch1, ch2)
	Process(merge)
	Consumer(Process(merge))
	waitGroup.Wait()
}
