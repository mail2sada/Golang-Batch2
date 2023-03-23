package main

import "fmt"

func main() {
	fmt.Println("Demo range loop")

	str := "Hello welcome to Mavenir"

	for idx, val := range str {

		fmt.Printf("\nstr[%d] value[%d] char[%c]", idx, val, val)
	}
}
