package main

import "fmt"

type Employee struct {
	EmpId int
	Name  string
}

//Print

func (a Employee) Print() {
	fmt.Println("======================")
	fmt.Println("EmpID:", a.EmpId)
	fmt.Println("Name:", a.Name)
	fmt.Println("------------------------")
}

//Update

func (a *Employee) UpdateName(name string) {
	a.Name = name
}

func (a Employee) Initialise(EmpId int, Name string) {
	a.EmpId = EmpId

	a.Name = Name
}

func main() {
	fmt.Println("Methods...")

	var e1 = Employee{EmpId: 1, Name: "Anilkumar"}
	var e2 = Employee{}
	e2.Initialise(10, "Test")

	e2.Print()
	e1.Print()

	e1.UpdateName("Vishal kumar")
	e1.Print()

}
