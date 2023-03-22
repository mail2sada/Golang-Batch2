// helloworld.go
// Few things to note
/*
1. Every go file has extension .go
2. every go file belongs to a package
*/

package main // this file belongs to main package.

import "fmt" // fmt package is imported, so that it functionalities can be used

// func is keyword, main is name of the function, () it is not taking any arguments
func main() {
	//fmt package has a function with the name Println
	// it can take any number of arguments and arguments can be of any type.
	fmt.Println("Welcome to Mavenir Go lang training")
}
