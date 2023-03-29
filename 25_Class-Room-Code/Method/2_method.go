package main

import "fmt"

type School struct {
	SchoolId int
	Name     string
	Address  string
}

type Office struct {
	OfficeName string
	Address    string
}

func (s School) getAddress() string {
	fmt.Println("In getAddress School")
	return s.Address
}

func (o Office) getAddress() string {
	fmt.Println("In get Address office")
	return o.Address
}

func main() {
	scl := School{SchoolId: 1, Name: "XYZ", Address: "Bangalore"}
	off := Office{OfficeName: "Mavenir", Address: "Manyata"}

	fmt.Println(scl.getAddress())
	fmt.Println(off.getAddress())
}
