package main

import "fmt"

func main() {
	fmt.Println("Demo: Array copy")
	var array1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	array2 := array1

	fmt.Println(array1)
	fmt.Println(array2)

	array2[5] = 500
	fmt.Println(array1)
	fmt.Println(array2)

	array3 := &array1
	fmt.Println(*array3)
	array1[0] = 10000
	fmt.Println(array1)
	fmt.Println(array3)
	array3[4] = 5000
	fmt.Println(array1)
	fmt.Println(array3)
}
