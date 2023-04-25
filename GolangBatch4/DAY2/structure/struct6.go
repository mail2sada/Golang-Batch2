package main

import "fmt"

type Student struct {
	roll_no int
	name    string
	class   string
}

func main() {
	fmt.Println("Demo: Heap allocation of structure")
	var std = Student{}
	std.roll_no = 1
	std.class = "X"
	std.name = "abc"
	std1 := new(Student)
	std1.roll_no = 2
	std1.name = "xyz"
	std1.class = "IX"
	fmt.Println(*std1)
	fmt.Println(std)
}
