package main

import "fmt"

func main() {
	fmt.Println("Demo: break")

	i := 20
	for {
		fmt.Println("Hello")
		i--
		if i == 15 {
			break
		}
	}

	i = 0

	for {
		i++
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
		if i == 100 {
			break
		}
	}
}
