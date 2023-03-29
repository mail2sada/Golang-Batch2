package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

/*
print emp details


-style=capital, lower, title
-format 2 digts, 3 dgits
-format =2
*/

type Emp struct {
	Empid     int
	firstName string
	lastName  string
	Salary    float32
}

var empSlice = []Emp{{Empid: 1, firstName: "abc", lastName: "def", Salary: 100000}, {Empid: 2, firstName: "xyz", lastName: "pqr", Salary: 50000}}

func main() {
	var (
		style  = flag.String("style", "capital", "style will output in capital, lower, title")
		format = flag.Int("format", 0, "It will print salary in specific order")
		
	)
	fmt.Println("Demo: Commandline tools")
	flag.Parse()

	

	fmt.Println("stype:", *style, "\nformat:", *format)

	switch *style {
	case "capital":
	case "lower":
	case "title":
	default:
		log.Fatalln("Invalid style can't proceed, allowed values are capital, upper, lower")
	}
	switch *format {
	case 1:
	case 2:
	default:
		log.Fatalln("Invalid format, allowed values are 1 or 2")
	}

	for _, val := range empSlice {

		fmt.Println("EmpId:", val.Empid)
		if *style == "capital" {
			fmt.Println("EmpId:", strings.ToUpper(val.firstName))
			fmt.Println("EmpId:", strings.ToUpper(val.lastName))
		} else if *style == "lower" {
			fmt.Println("EmpId:", strings.ToLower(val.firstName))
			fmt.Println("EmpId:", strings.ToLower(val.lastName))
		} else {
			fmt.Println("EmpId:", strings.ToTitle(val.firstName))
			fmt.Println("EmpId:", strings.ToTitle(val.lastName))
		}

		if *format == 1 {
			fmt.Printf("\n%0.1f\n", val.Salary)
		} else {
			fmt.Printf("\n%0.2f\n", val.Salary)

		}
	}
}
