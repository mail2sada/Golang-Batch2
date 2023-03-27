package main

import "fmt"

func main() {
	fmt.Println("Demo: Elipsis")

	var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(array)

	var myTestArray = [...]int{100: 500, 50: 100, 500: -1}

	fmt.Println(myTestArray)
}
