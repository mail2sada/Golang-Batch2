package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type MyInterface interface {
	String() string
}

type Student struct {
	RollNo int
	Name   string
	Class  string
}

func (std Student) String() string {
	ret := strconv.Itoa(std.RollNo)
	ret += "\n" + std.Name
	ret += "\n" + std.Class
	return ret
}

func main() {
	fmt.Println("Demo: Multiple interface")

	var std MyInterface = Student{RollNo: 1, Name: "Dilip", Class: "Mtech"}

	ret := std.String()
	fmt.Println(reflect.TypeOf(ret))
	fmt.Println(std)
}
