//
//
//generator->squres -> display
//			  squre
// generator->squre  display
//			  squre

package main

import (
	"fmt"
	"sync"
)

const FanOutConfig = 5

func Generate(num ...int) chan int {
	//wg.Add(1)
	gen := make(chan int)
	go func() {
		//defer wg.Done()
		defer close(gen)
		for _, val := range num {
			gen <- val
		}
	}()
	return gen
}

func Squre(ch chan int) chan int {
	//wg.Add(1)
	sqr := make(chan int)
	swg := sync.WaitGroup{}

	for i := 0; i < FanOutConfig; i++ {
		swg.Add(1)
		go func() {
			defer swg.Done()
			for val := range ch {
				squre := val * val
				sqr <- squre
			}
		}()
	}
	go func() {
		swg.Wait()
		close(sqr)
		//defer wg.Done()
	}()
	return sqr
}

func Display(ch chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println(val)
		}
	}()
}

var wg = sync.WaitGroup{}

func main() {

	// gch := Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// sch := Squre(gch)
	// Display(sch)

	Display(Squre(Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)))
	wg.Wait()

}
