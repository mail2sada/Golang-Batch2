package main

import "fmt"

func main() {
	fmt.Println("Demo: Annonynous struct")

	a := struct {
		header int8
		body   []byte
	}{header: 123, body: []byte("Hello")}

	fmt.Println(a)
}
