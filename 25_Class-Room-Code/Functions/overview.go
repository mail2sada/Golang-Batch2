package main

import (
	"fmt"
	"strings"
)

func main() {
	var greeetingList = []string{"Hello...", "Welcome", "Namaste..."}

	for _, val := range greeetingList {
		toUpper(val)
		fmt.Println(val)
		toUpperByRef(&val)
		fmt.Println(val)
	}

	fmt.Println(greeetingList)
}

func toUpperByRef(str *string) {
	*str = strings.ToUpper(*str)
	//return str
}

func toUpper(str string) {
	str = strings.ToUpper(str)
	//return str
}

// func displayGreetingWithArg(greeting string) {
// 	fmt.Println(toUpper(greeting))
// }

// func displayGreeting() {
// 	fmt.Println("Hello everyone...")
// }


