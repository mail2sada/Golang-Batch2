package main

import "fmt"

type employee struct {
	empId int
	name  string
	dept  string
	role  string
	age   int
}

func main() {
	fmt.Println("Demo: Struct")

	var emp1 employee = employee{empId: 123, name: "Anil kumar", dept: "MDE", role: "Engineer", age: 23}

	fmt.Println(emp1.empId)
	fmt.Println(emp1.name)
	fmt.Println(emp1.dept)
	fmt.Println(emp1.age)
	fmt.Println(emp1.role)

	emp1.empId = 345

	fmt.Println(emp1)

}
