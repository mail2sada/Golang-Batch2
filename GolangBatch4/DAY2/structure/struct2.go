package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
	emilId  string
}

type Employee struct {
	EmpID           int
	Role            string
	Dept            string
	PersonalDetails Person
}

type HealthClub struct {
	memberID        int
	PersonalDetails Person
}

func main() {
	fmt.Println("Demo nested structures")

	p1 := Person{name: "Anand", age: 32, address: "HSR", emilId: "anand@mavenir.com"}
	p2 := Person{name: "Vijay", age: 36, address: "BTM", emilId: "vijay@mavenir.com"}

	emp1 := Employee{EmpID: 1, Role: "Manager", Dept: "NGN", PersonalDetails: p1}

	emp2 := Employee{EmpID: 2, Role: "Architech", Dept: "Radio", PersonalDetails: p2}

	mem1 := HealthClub{memberID: 1, PersonalDetails: p1}
	mem2 := HealthClub{memberID: 2, PersonalDetails: p2}

	fmt.Println(emp1)
	fmt.Println(emp2)

	fmt.Println(mem1)
	fmt.Println(mem2)
}
