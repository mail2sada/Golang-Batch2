package main

import "fmt"

type Employee struct {
	int
	string
	dept    string
	address string
}

func main() {
	emp := Employee{int: 12, string: "Anand Rao", dept: "Radio", address: "Banglore"}

	ptrEmp := &emp

	fmt.Println(*ptrEmp)

	ptrEmp.address = "HSR Bangalore"

	fmt.Println(*ptrEmp)

}
