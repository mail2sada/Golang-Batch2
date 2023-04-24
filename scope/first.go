package main

import "fmt"

var global int

func main() {
	var a = 100
	fmt.Println(a)
	{
		var b = 1000
		fmt.Println("Value of b", b)
	}

	fmt.Println("Demo:Scope...")
	fmt.Println("Value of global is", global)
	test()
	fmt.Println("Value of global after calling test", global)
	myTest()

	fmt.Println("Value of global after calling mytest", global)
}

func test() {
	global = 100
}
