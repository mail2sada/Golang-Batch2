package main

import "fmt"

func main() {

	var s1, s2, s3, s4, s5 int
	var sum int
	var avg float64

	fmt.Println("Enter marks in 5 subjects")

	fmt.Scanf("%d", &s1)
	fmt.Scanf("%d", &s2)
	fmt.Scanf("%d", &s3)
	fmt.Scanf("%d", &s4)
	fmt.Scanf("%d", &s5)

	sum = s1 + s2 + s3 + s4 + s5

	avg = float64(sum) / float64(5)

	fmt.Println("Average marks ", avg)

}
