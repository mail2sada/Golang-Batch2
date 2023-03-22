package main

import "fmt"

func main() {
	fmt.Println("Demo: declarining multiple variables")

	var v1, v2, v3 uint = 10, 20, 30

	fmt.Println("Value of v1 is", v1, "\nValure of v2 ", v2, "\nValue of v3 ", v3)

	var t1, t2, t3 int

	t1 = 10

	t2 = 20

	t3 = -100

	fmt.Println("Value of t1 is", t1, "\nValure of t2 ", t2, "\nValue of t3 ", t3)

	var myInt, myFloat, myString, myBool = 55, 36.82, "Hello everyone", true

	fmt.Printf("Type of myInt %T\nType of myFloat %T\nType of myString %T\nType of myBool %T", myInt, myFloat, myString, myBool)

	fmt.Println("Value of myInt, myFloat, myString and myBool are", myInt, myFloat, myString, myBool)

}
