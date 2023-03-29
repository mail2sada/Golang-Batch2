package main

import "fmt"

type address struct {
	building_no string
	area        string
	city        string
	pincode     int
}

type employee struct {
	empId         int
	firstName     string
	lastName      string
	dept          string
	addr          address
	work_location string
}

func main() {

	fmt.Println("Demo: Structure")

	var addr address = address{building_no: "123", area: "Nagawara", city: "Bangalore", pincode: 560045}

	var emp1 employee = employee{empId: 12345,
		firstName:     "Dilip",
		lastName:      "Kumar",
		dept:          "MDE",
		addr:          address{building_no: "123", area: "HSR", city: "Bangalore", pincode: 560102},
		work_location: "Bangalore"}

	fmt.Println(addr)
	fmt.Println(emp1)

	emp1.dept = "NGN"

	fmt.Println(emp1.firstName, emp1.lastName, emp1.empId, emp1.dept, emp1.work_location)
}
