package main

import "fmt"

func main() {
	fmt.Println("Demo: Function with return")
	fmt.Println("Value of ", Add(10, 50))

	Variadic("Hello", "HI", "Test")

	fmt.Println("Result is ", VariadicAdd(10, 20, 30, 40, 50, 60, 10))

	fmt.Println("Result is ", VariadicAdd(10, 20, 10))

}

func Add(x, y int) int {
	fmt.Println("Inside add")
	c := x + y
	return c
}

func Variadic(str string, elements ...string) {
	fmt.Println("Variadic:", str)
	for _, vak := range elements {
		fmt.Println(vak)
	}
}

func VariadicAdd(intgers ...int) int {
	sum := 0
	for _, val := range intgers {
		sum += val
	}
	return sum
}

func ADD(x int, y int) {

}
