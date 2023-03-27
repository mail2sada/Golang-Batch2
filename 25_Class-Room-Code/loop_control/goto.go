package main

import "fmt"

func main() {
	fmt.Println("Go to Statement")

Label1:
	i := 0
	fmt.Println("I am here..")
	for {
		i++
		if i%2 == 0 {
			goto Label1
		}
		fmt.Println("Value of i", i)
		if i == 10 {
			break
		}
	}
}
