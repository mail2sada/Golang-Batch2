package main

import "fmt"

func main() {

	fmt.Println("Demo: Arrays")

	//var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var array [10]int = [10]int{0: 1, 5: 10, 9: 100}

	var myArray = [10]string{2: "Mavenir", 5: "Test", 7: "Code"}

	fmt.Println(array)
	fmt.Println(myArray)

	sdoArray := [3]int{1, 2, 3}

	fmt.Println(sdoArray)

	array[0] = 100

	i := array[1]

	sum := 0

	for i := 0; i < 10; i++ {
		sum += array[i]
	}

	fmt.Println("Sum is", sum)
}
