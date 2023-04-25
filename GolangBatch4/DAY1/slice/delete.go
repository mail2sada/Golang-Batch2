package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice delete...")

	var testSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	myfrontslice := testSlice[:5]
	fmt.Println(myfrontslice)

	myfrontslice[3] = 100
	fmt.Println(myfrontslice)
	fmt.Println(testSlice)
	mytailslice := testSlice[7:]
	fmt.Println(mytailslice)
	middleSlice := testSlice[3:7]
	fmt.Println(middleSlice)

	testSlice = append(testSlice[:5], testSlice[6:]...)
	fmt.Println(testSlice)
	

}
