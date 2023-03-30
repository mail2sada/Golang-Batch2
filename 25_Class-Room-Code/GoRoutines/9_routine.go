/*

Producer
			Muxer Processor Outpur
Proceur

*/

package main

import (
	"fmt"
	"sync"
)

func Producer(numbers ...int) chan int {
	channel := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(channel)
		fmt.Println("In Procuder Routine...")
		for _, val := range numbers {
			channel <- val
		}
	}()
	return channel
}

func Muxer(ch1 <-chan int, ch2 <-chan int) chan int {
	muxer := make(chan int)
	wg.Add(1)
	go func() {
		var ch1open, ch2open bool = true, true
		defer wg.Done()
		defer close(muxer)
		i := 0
		for i < 10 {
			select {
			case val, ch1open := <-ch1:

				if ch1open {
					muxer <- val
				}

			case val, ch2open := <-ch2:

				if ch2open {
					muxer <- val
				}
			}
			if ch2open == false && ch1open == false {
				break
			}
			i++
			fmt.Println("Debug:::", ch1open, ch2open)
		}
	}()

	return muxer
}

func Processor(in <-chan int) chan int {
	process := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(process)
		fmt.Println("In Processor...")
		for val := range in {
			process <- val * val
		}
	}()
	return process
}

func Consumer(in <-chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range in {
			fmt.Println(val)
		}
	}()
}

var wg = sync.WaitGroup{}

func main() {

	fmt.Println("Demo: FanIn")

	ch1 := Producer(1, 3, 5, 7, 9, 11, 13)

	ch2 := Producer(2, 4, 6, 8, 10, 12, 14, 16)

	mux := Muxer(ch1, ch2)

	proc := Processor(mux)

	Consumer(proc)

	wg.Wait()
}
