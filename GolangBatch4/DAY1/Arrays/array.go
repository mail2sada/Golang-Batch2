package main

import "fmt"

func main() {

	var firstArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println(myArray)

	myArray[0] = 100
	a := myArray[4]
	fmt.Println("Value of a is ", a)
	fmt.Println(myArray)
	fmt.Println(firstArray)

	testArray := [5]int{1: 100, 3: 200}
	fmt.Println(testArray)

	mySecondArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySecondArray)

	var array = [...]int{0: 100, 5, 102: 100, 10: 200, 100: 10000, 58: 100000}
	fmt.Println(array, len(array))

	for i := 0; i < len(array); i++ {
		fmt.Println("Value of array[", i, "] is", array[i])
		array[i] = 100
	}

	for idx, val := range array {
		fmt.Println("value at ", idx, "is", val)
		val = 1000
	}

	for i := 0; i < len(array); i++ {
		fmt.Println("Value of array[", i, "] is", array[i])
	}
}
