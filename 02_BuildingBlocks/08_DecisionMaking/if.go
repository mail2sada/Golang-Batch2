package main

import "fmt"

func main() {
	fmt.Println("Demo if statement")

	const (
		Sunday int = iota
		Monday
		Tuesday
		Wednsday
		Thursday
		Friday
		Saturday
	)

	weekDay := Monday

	if weekDay == Sunday {
		fmt.Println("Sunday")
	}

	if weekDay == Monday {
		fmt.Println("Monday")
	}

	if weekDay == Tuesday {
		fmt.Println("Tuesday")
	}

	if weekDay == Wednsday {
		fmt.Println("Wednsday")
	}

	if weekDay == Thursday {
		fmt.Println("Thursday")
	}

	if weekDay == Friday {
		fmt.Println("Friday")
	}

	if weekDay == Saturday {
		fmt.Println("Saturday")
	}
}
