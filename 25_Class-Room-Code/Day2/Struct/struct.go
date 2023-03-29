package main

import "fmt"

type employee struct {
	empId         int
	firstName     string
	lastName      string
	dept          string
	address       string
	work_location string
}

func main() {

	fmt.Println("Demo: Structure")

	var emp1 employee = employee{empId: 12345,
		firstName:     "Dilip",
		lastName:      "Kumar",
		dept:          "MDE",
		address:       "Bangalore",
		work_location: "Bangalore"}

	fmt.Println(emp1)

	emp1.dept = "NGN"

	fmt.Println(emp1.firstName, emp1.lastName, emp1.empId, emp1.dept, emp1.work_location)
}
