package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Demo: String functions...")

	str1 := "Hello world...§"

	if strings.Contains(str1, "H") {
		fmt.Println("Contains H")
	}

	if strings.ContainsAny(str1, "eloa") {
		fmt.Println("Available")
	}
	if strings.ContainsRune(str1, '§') {
		fmt.Println("Contains §")
	}

	str1 = strings.ToUpper(str1)
	fmt.Println(str1)

	str1 = strings.ToLower(str1)
	fmt.Println(str1)

	str1 = strings.ToTitle(str1)
	fmt.Println(str1)

}
