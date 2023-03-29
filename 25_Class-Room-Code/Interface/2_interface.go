package main

import (
	"fmt"
	"strconv"
)

type School struct {
	SchoolName string
	Address
}

func (sch School) String() string {
	ret := "School Name:" + sch.SchoolName
	ret += fmt.Sprintln(sch.Address)
	return ret
}

type Office struct {
	OfficeName string
	Address
}

func (off Office) String() string {
	ret := "Office Name:" + off.OfficeName
	ret += fmt.Sprintln(off.Address)
	return ret
}

type Home struct {
	HomeName string
	Address
}

func (home Home) String() string {
	ret := "Office Name:" + home.HomeName
	ret += fmt.Sprintln(home.Address)
	return ret
}

type Address struct {
	DoorNo  string
	Area    string
	City    string
	Pincode int
}

func (addr Address) String() string {
	ret := "Door No:" + addr.DoorNo
	ret += "\nArea:" + addr.Area
	ret += "\nCity:" + addr.City
	ret += "\nPincode" + strconv.Itoa(addr.Pincode)
	return ret
}

func main() {
	off := Office{OfficeName: "Mavenir", Address: Address{DoorNo: "MFAR 8", Area: "Manyata", City: "Bangalore", Pincode: 560045}}

	fmt.Println(off)
}
