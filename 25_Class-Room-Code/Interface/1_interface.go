package main

import (
	"fmt"
	"strconv"
)

type iPrint interface {
	Print()
}

type School struct {
	SchoolId int
	Name     string
}

func (sch School) Print() {
	fmt.Println(sch.SchoolId)
	fmt.Println(sch.Name)
}

func (sch School) String() string {
	fmt.Println("In String of sch")
	ret := "School ID is " + strconv.Itoa(sch.SchoolId) + "\n"
	ret += "School Name is" + sch.Name

	return ret
}

type Office struct {
	OfficeName string
}

func (off Office) Print() {
	fmt.Println(off.OfficeName)
}

func main() {

	// var v1, v2 iPrint

	sch := School{SchoolId: 1, Name: "Abc"}

	// off := Office{OfficeName: "Mavenir"}

	// v1 = sch

	// v2 = off

	// v1.Print()

	// v2.Print()

	fmt.Println(sch)

	str := fmt.Sprintf("%s", sch)

	fmt.Println(str)

}
