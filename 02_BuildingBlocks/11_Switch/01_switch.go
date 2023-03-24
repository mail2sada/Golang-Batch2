package main

import "fmt"

type Week uint8

const (
	Sunday Week = iota
	Monday
	Tuesday
	Wednsday
	Thursaday
	Friday
	Saturday
)

func main() {
	fmt.Println("Demo: Switch statement")

	dayOFWeek := Saturday

	switch dayOFWeek {
	case Sunday:
		fmt.Println("Its Sunday, last day of your weekend")
	case Monday:
		fmt.Println("Its Monday, beginning of the week")
	case Tuesday:
		fmt.Println("Its Tuesday...")
	case Wednsday:
		fmt.Println("Its Wednsday, I have lot of work")
	case Thursaday:
		fmt.Println("Its Thurday, I am about to finish my work")
	case Friday:
		fmt.Println("Finally Its Friday, weekend will start..")
	case Saturday:
		fmt.Println("OhHo!! Its Saturday!! fun time...")
	default:
		fmt.Println("I dont know who you are")
	}

}
