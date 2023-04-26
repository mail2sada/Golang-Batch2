package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Print() {
	fmt.Println("Person details are")
	fmt.Println("Name", p.name)
	fmt.Println("age", p.age)
	fmt.Println("address", p.address)
}

type Address struct {
	address string
}

func (a Address) Print() {
	fmt.Println(a.address)
}

type Employee struct {
	eid         int
	dept        string
	designation string
	Person
	a Address
}

func main() {
	fmt.Println("Prompted methods")
	e1 := Employee{eid: 1, dept: "NGN", designation: "Eng", Person: Person{name: "Anand", age: 35, address: "Bangalore"}}
	e1.Print()
	e1.a.Print()
}
