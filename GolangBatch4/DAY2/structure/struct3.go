package main

import "fmt"

func main() {
	type Person struct {
		a int
		b string
	}
	var p Person = Person{a: 10, b: "hello"}
	fmt.Println(p)
}
