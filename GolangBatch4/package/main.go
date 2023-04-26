package main

import (
	"fmt"

	cl "image/color"
	math "tutorial/maths"
	"tutorial/maths/advanced"
	"tutorial/test"

	c1 "github.com/fatih/color"
)

func init() {
	c1.Red("This is 1")
	fmt.Println("1 Init inside main function ")
}

func init() {
	a := cl.Gray16Model
	fmt.Println(a)
	c1.Green("This is 2")
	fmt.Println("2 Init inside main function ")
}

func main() {
	c1.Yellow("Hi")
	a := math.Add(10, 20)
	fmt.Println(a)
	b := math.Sub(20, 10)
	fmt.Println(b)
	c := advanced.Square(10)

	d := advanced.Sin(90)
	fmt.Println(c, d)
	test.TestCode()

}
