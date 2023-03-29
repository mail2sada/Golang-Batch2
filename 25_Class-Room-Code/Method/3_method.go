package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	Building string
	Area     string
	City     string
	Pincode  int
}

func (addr Address) PrintAddress() {
	fmt.Println(addr.Building)
	fmt.Println(addr.Area)
	fmt.Println(addr.City)
	fmt.Println(addr.Pincode)
}

type School struct {
	SchoolId int
	Name     string
	Address
}

type Office struct {
	OfficeName string
	Address
}

func main() {

	schoolAddr := Address{Building: "A101", Area: "HSR", City: "Bangalore", Pincode: 560102}
	officeAddr := Address{Building: "B102", Area: "Indiranagar", City: "Bangalore", Pincode: 560049}
	scl := School{SchoolId: 1, Name: "XYZ", Address: schoolAddr}
	off := Office{OfficeName: "Mavenir", Address: officeAddr}

	office := new(Office)

	fmt.Println(reflect.TypeOf(office))

	scl.PrintAddress()

	off.PrintAddress()

	off.Building = "T102"

	fmt.Println(off.Building)
	fmt.Println(scl.Area)

	off.PrintAddress()

}
