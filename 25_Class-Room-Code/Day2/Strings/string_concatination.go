package main

import "fmt"

func main() {

	fmt.Println("Demo: String concatination")

	str1 := "Hello"

	str2 := " how are you?"

	str4 := "I am good"

	str3 := ""

	str3 = str1 + str2 + str4

	fmt.Println(str3)

	str1 += str2

	fmt.Println(str1)

	str1 += str4

	fmt.Println(str1)

	fmt.Printf("%c\n", str1[0])

	

}
