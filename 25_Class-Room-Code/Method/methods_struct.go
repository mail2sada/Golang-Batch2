package main

import "fmt"

type Emp struct {
	empId string

	name string
}

func (e Emp) print() {
	fmt.Println(e.empId)
	fmt.Println(e.name)
}

func (e Emp) isEqual(b Emp) bool {
	if e == b {
		return true
	}
	return false
}

func main() {
	fmt.Println("Demo...")
	var a, b = Emp{"123", "ABC"}, Emp{"456", "DEF"}

	a.print()
	b.print()

	if a.isEqual(b) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal..")
	}

}
