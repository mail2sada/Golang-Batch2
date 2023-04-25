package main

import "fmt"

func main() {
	fmt.Println("Demo: Operators")

	var a, b = 10, 20

	sum := a + b
	dif := a - b
	pro := a * b
	qua := a / b
	rem := a % b

	fmt.Println(sum, dif, pro, qua, rem)

	// Logical operators

	// &&

	// ||

	res := a == 10 || b != 20
	fmt.Println(res)

	res = a == 10 && b != 20
	fmt.Println(res)

	res = !(a == 10)

	fmt.Println(res)

	// relational operators
	// <
	// >
	// <=
	// >=
	// ==
	// !=

	//assignment operators
	// a +=b // a = a +b
	// = 
	// +=
	// *=
	// -=
	// &
	// |
	// <<
	// >>
	

}
