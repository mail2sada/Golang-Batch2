package main

import "fmt"

type Employee struct {
	EmpID     int
	FirstName string
	LastName  string
	Dept      string
}

func (emp Employee) print() {
	fmt.Println("--------------------------")
	fmt.Println("EMP ID\t FirstName\t Lastname \t Dept")
	fmt.Println("---------------------------")
	fmt.Println(emp.EmpID, "\t", emp.FirstName, "\t", emp.LastName, "\t", emp.Dept)
}

func (emp *Employee) Initailize(eid int, fname string, lname string, dept string) {
	emp.EmpID = eid
	emp.FirstName = fname
	emp.LastName = lname
	emp.Dept = dept
}

func main() {
	//emp := Employee{EmpID: 1, FirstName: "Anil", LastName: "Kumar", Dept: "RCS"}
	emp := Employee{}

	emp.Initailize(1, "Anand", "D", "NGN")

	emp.print()
}
