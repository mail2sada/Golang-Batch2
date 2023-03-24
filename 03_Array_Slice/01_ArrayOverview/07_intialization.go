package main

import "fmt"

func main() {

	fmt.Println("Demo: Declaration and initialisation with elipsis")

	var arrayInt = [...]int{10, 20, 30, 40, 50}

	fmt.Println("Contents of arrayInt:", arrayInt)

	strArray := [...]string{"Red", "Green", "Blue"}

	fmt.Println("Contents of strArray:", strArray)

	strNewArray := [...]string{1: "Hello", 5: "Namaste", 10: "How are you?"}

	fmt.Println("Contents of strNewArray:", strNewArray)
}
