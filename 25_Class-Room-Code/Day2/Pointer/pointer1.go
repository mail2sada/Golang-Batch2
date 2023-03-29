package main

import "fmt"

func main() {
	var str = "Hello how are you"

	var ptrStr *string = &str

	*ptrStr = "abcde"

	
}
