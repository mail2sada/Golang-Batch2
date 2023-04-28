package main

import "fmt"

type CustomType struct {
	a int
}

func (c CustomType) String() string {
	return fmt.Sprintf("Value of \"a\":%d", c.a)
}

func main() {

	fmt.Println("Interfaces")

	a := CustomType{a: 100}
	fmt.Print(a)
	b := CustomType{a: 10000}
	fmt.Print(b)
}
