package main

import "fmt"

func main() {
	fmt.Println("Reading from key board")
	var a, b, c int
	var flt float64
	fmt.Println("Please enter values for a,b,c")

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%f", &flt)

	fmt.Println(a, b, c)
	fmt.Println(flt)
}
