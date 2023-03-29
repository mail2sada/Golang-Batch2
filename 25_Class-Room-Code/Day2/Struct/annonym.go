package main

import "fmt"

var ab = struct {
	name string
	dob  string
}{
	name: "Anil Kumar",
	dob:  "10-07-1998",
}

func main() {
	fmt.Println("Demo: Annonymous struct")

	a := struct {
		name string
		dob  string
	}{
		name: "Anil Kumar",
		dob:  "10-07-1998",
	}

	fmt.Println(a)
	fmt.Println(ab)
}
