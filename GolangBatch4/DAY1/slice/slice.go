package main

import "fmt"

func main() {
	var intSlice []int = []int{1, 2, 10: 3, 4, 5, 6}

	fmt.Println(intSlice)

	fmt.Println(len(intSlice))

	intSlice = append(intSlice, 10, 11, 12, 13, 14, 15)
	fmt.Println("len", len(intSlice))
	fmt.Println(intSlice)

	fmt.Println("cap:", cap(intSlice))

	intSlice = append(intSlice, 101, 102, 103, 104, 105, 106, 107, 108)

	fmt.Println("len", len(intSlice))
	fmt.Println("cap:", cap(intSlice))
	p := &intSlice[0]
	fmt.Println(p)
	intSlice = append(intSlice, 10)
	p = &intSlice[0]
	fmt.Println(p)
	fmt.Println("len", len(intSlice))
	fmt.Println("cap:", cap(intSlice))

	mySlice := make([]int, 3, 5)
	fmt.Println(mySlice)
	fmt.Println("Cap", cap(mySlice))
	fmt.Println("Len", len(mySlice))
	mySlice = append(mySlice, 10, 11, 12, 13, 14)

	fmt.Println(mySlice)
	fmt.Println("Cap", cap(mySlice))
	fmt.Println("Len", len(mySlice))

	
}
