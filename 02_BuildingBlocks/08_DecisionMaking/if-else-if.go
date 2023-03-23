package main

import "fmt"

func main() {
	fmt.Println("Demo: if else if")

	const (
		Sunday int = iota
		Monday
		Tuesday
		Wednsday
		Thursday
		Friday
		Saturday
	)

	weekday := Friday

	if weekday == Sunday {
		fmt.Println("Weekday is Sunday")
	} else if weekday == Monday {
		fmt.Println("Weekday is Monday")
	} else if weekday == Tuesday {
		fmt.Println("Weekday is Tuesday")

	} else if weekday == Wednsday {
		fmt.Println("Weekday is Wednsday")

	} else if weekday == Thursday {
		fmt.Println("Weekday is Thursday")

	} else if weekday == Friday {
		fmt.Println("Weekday is Friday")

	} else if weekday == Saturday {
		fmt.Println("Weekday is Saturday")
	}
}
