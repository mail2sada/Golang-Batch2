package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Demo: String methods")

	var str = "helloabcthisabcisabcgoabctraining"

	splitstr := strings.Split(str, "abc")

	for _, val := range splitstr {
		fmt.Println(val)
	}
	joinstr := strings.Join(splitstr, "/*/")
	fmt.Println(joinstr)
	fmt.Println("Length of joinstr", len(joinstr))
}
