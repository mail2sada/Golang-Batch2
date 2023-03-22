package main

import "fmt"

// this variable is accessible through out main package.
// this is not exported variable hence it can not be accessed outside package
// more explaination can be found in Modules and Packages Section
var global int = 100

func main() {

	fmt.Println("Demo scope of variable")

	fmt.Println("Value of global variable is ", global)
	var localToMain = 500
	{
		fmt.Println("Value of localToMain is ", localToMain)
		fmt.Println("Value of the global int is accessible inside block", global)

		var localToBlock = 1000

		fmt.Println("localToBlocl variable is accesible only insid this block ", localToBlock)

	}

	fmt.Println("localToMain and gloabl are accessible here ", localToMain, global)

	fmt.Println("Scope of localToBlock is limited only to scope any attempt to access this here will result in compiler error")

	fmt.Println("We have one more go file by name as helper.go, we have defined a string globally ")

	fmt.Println("Variable declared in helper.go is accessible here", globalGreeting)
}
