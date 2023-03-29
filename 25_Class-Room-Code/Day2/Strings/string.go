package main

import "fmt"

func main() {

	fmt.Println("Demo: Strings...")

	var str = "Hello \n Mavenir \\n"

	var backtickstr = `hello how are you?
		prining contents of string
		""" \nakjhsjdh as aksjhdas`

	fmt.Println(str)

	fmt.Println(backtickstr)

	fmt.Println(str[0])
}
