package main

import (
	"fmt"
)

func main() {
	fmt.Println("Demo: String package")
	str := `Hello this is golang code to demo strings package!@#$#$%^FGGH`

	fmt.Println("str:", str)

	str2 := "This is Golang training"

	fmt.Println("str2:", str2)

	str3 := str + str2
	fmt.Println("str3", str3)

	str3 += "This go program strings1.go"

	fmt.Println(str3)

}
