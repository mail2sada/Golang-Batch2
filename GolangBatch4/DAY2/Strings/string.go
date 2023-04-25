package main

import "fmt"

func main() {

	var str = "Hello and welcome to Golang training"
	fmt.Println(str)
	var newStr = `This is my string \n """"
	skkjksdjks s
	skdjskjsd
	skjdskjd
	hellow`

	fmt.Println("This is my new string ", newStr)

	for idx, val := range newStr {
		//fmt.Println(idx, ":", val)
		fmt.Printf("%d:%c\n", idx, val)
	}

	val1 := newStr[0]

	fmt.Println(val1)
	newStr[0] = val1
}
