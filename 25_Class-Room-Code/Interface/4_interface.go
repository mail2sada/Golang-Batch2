package main

import "fmt"

type Intf1 interface {
	func1()
	func2()
}

type Intf2 interface {
	func1()
}

type Intf3 interface {
	func2()
}

type Data struct {
	element string
}

func (d Data) func1() {
	fmt.Println("In func1")
}

func (d Data) func2() {
	fmt.Println("In func2")
}

func main() {
	fmt.Println("Multiple interfaces")

	var intf1 Intf1 = Data{element: "Hello how are you"}

	intf1.func1()
	intf1.func2()

	var intf2 Intf2 = intf1
	intf2.func1()
	fmt.Println(intf2)

	var intf3 Intf3 = intf1
	intf3.func2()
}
