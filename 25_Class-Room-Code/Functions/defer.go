package main

import "fmt"

func main() {
	fmt.Println("Demo: defer...")

	a := 100

	ptr := &a

	defer func(x *int) {
		fmt.Println(*x)
	}(ptr)

	a = 10000



	myFunc()

	defer func () {
		
	}()
	defer fmt.Println("I want to defer this print")

	defer fmt.Println("Defer this statement as well")

	fmt.Println("Inside main")

	fmt.Println("Exiting main")
}

func myFunc() {

	fmt.Println("In myFunc")

	defer fmt.Println("In myfunc: I want to defer this")

	fmt.Println("Exiting myFunc")

}
