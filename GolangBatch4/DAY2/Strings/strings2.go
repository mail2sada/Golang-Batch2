package main

import "fmt"

func main() {
	fmt.Println("Demo: Strings with fmt")

	a := 10

	pi := 3.142

	str := fmt.Sprintf("Value of pi is %f and a is %d", pi, a)

	str2 := fmt.Sprint("Value of a is ", a, " and pi is ", pi)

	fmt.Println(str)
	fmt.Println("Str2:", str2)

}
