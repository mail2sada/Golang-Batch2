package main

import "fmt"

// func CircleProperies(radius float64) (float64, float64) {
// 	const pi = 3.14159
// 	area := pi * radius * radius
// 	circ := 2 * pi * radius

// 	return area, circ

// }

func CircleProperies(radius float64) (area float64, circ float64) {
	const pi = 3.14159
	area = pi * radius * radius
	circ = 2 * pi * radius

	return area, circ

}

func main() {
	fmt.Println("Demo: Function returning multiple values")

	fmt.Println("We are testing function returning multiple values")

	_, c := CircleProperies(3.0)

	fmt.Println("Area is", a)
	//fmt.Println("Circufrence is", c)
}
