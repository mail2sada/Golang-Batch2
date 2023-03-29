package main

import "fmt"

type Base struct {
	color string
}

func (b Base) Say() {
	fmt.Println("Say from base is called...")
}

type Derived struct {
	Base
	style string
}

func main() {
	base := Base{color: "Red"}

	derived := Derived{Base: base, style: "My style"}

	derived.Say()

	fmt.Println("Hello", derived.color)

}

func (str *string) myString() {

}
