package main

import "fmt"

func main() {
	fmt.Println("Demo: Type conversion")

	var physics, chemestry, maths, english, hindi int = 99, 96, 100, 97, 80

	fmt.Println("Marks in Physics:", physics)
	fmt.Println("Marks in Chemestry:", chemestry)
	fmt.Println("Marks in Maths:", maths)
	fmt.Println("Marks in English:", english)
	fmt.Println("Marks in hindi:", hindi)

	var subject = 5

	var totalMarks = physics + chemestry + maths + english + hindi

	var average float64 = float64(totalMarks) / float64(subject)

	fmt.Println("Average score is ", average)
}
