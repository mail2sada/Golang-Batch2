package main

import (
	"fmt"
	"sync"
)

func main() {

	var slice = []string{"Hello", "Good Morning", "This is my test code", "This is a wonderful day", "Great to learn goroutin"}
	wg := sync.WaitGroup{}
	wg.Add(len(slice))
	for idx, val := range slice {
		go func(str string, cnt int) {
			defer wg.Done()

			for _, v := range str {
				fmt.Printf("%d:%c\n", cnt, v)
			}
			fmt.Println("idx:", cnt, val)
		}(val, idx)
	}
	wg.Wait()
}
