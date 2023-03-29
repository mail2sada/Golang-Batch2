package main

import "fmt"

func main() {

	fmt.Println("Demo: Sprintf")

	str1 := "Mavenir.com"
	str2 := "NGN"
	str3 := "MDE"

	str4 := fmt.Sprintf("%s/%s/%s", str1, str2, str3)

	fmt.Println(str4)

	str4 = fmt.Sprint("Hello\n", "Hi\n", "How are you\n")

	fmt.Println(str4)

}
