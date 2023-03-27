package main

import "fmt"

const (
	Sunday int = iota
	Monday
	Tuesday
	Wednsday
	Thursaday
	Friday
	Saturday
)

func main() {

	fmt.Println("Demo: Switch...")

	var WeekDay int

	fmt.Println("Enter weekday")

	fmt.Scanf("%d", &WeekDay)

	switch WeekDay {
	case Sunday:
		fmt.Println("Last day of weekend")
		fallthrough
	case Monday:
		fmt.Println("1st working day of week")
	case Tuesday:
		fmt.Println("1")
	case Wednsday:
		fmt.Println("2")
	case Thursaday:
		fmt.Println("3")
	case Friday:
		fmt.Println("4")
	case Saturday:
		fmt.Println("5")
	default:
		fmt.Println("Unhandled")
	}

	switch WeekDay {
	case Monday, Tuesday, Wednsday, Thursaday, Friday:
		fmt.Println("Weekdays")
	case Saturday, Sunday:
		fmt.Println("Weekends")

	}

}
