// ReadData
// ProcessData
// Output Data

// ReadData -> ProcessData --> Output pipeline

// generate -> squre -> output

package main

import (
	"fmt"
	"sync"
)

func Generate(num ...int) (gen chan int) {
	gen = make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(gen)
		for _, val := range num {
			gen <- val
		}
	}()
	return
}

func Squre(ch chan int) chan int {
	sqr := make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(sqr)
		for val := range ch {
			squre := val * val
			sqr <- squre
		}
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
	//wg.Add(4)
	// gch := Generate(1, 2, 3, 4, 5, 6, 7, 8, 10)
	// sch := Squre(gch)
	// Display(sch)

	Display(Squre(Squre(Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))))
	wg.Wait()
}
