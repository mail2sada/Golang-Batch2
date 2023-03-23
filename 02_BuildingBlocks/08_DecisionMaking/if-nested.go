package main

import "fmt"

func main() {

	fmt.Println("Demo: Nested if")

	const (
		Sunday int = iota
		Monday
		Tuesday
		Wednsday
		Thursday
		Friday
		Saturday
	)

	weekDay := Tuesday

	if weekDay != Sunday {
		if weekDay != Monday {
			if weekDay != Tuesday {
				if weekDay != Wednsday {
					if weekDay != Thursday {
						if weekDay != Friday {
							if weekDay != Saturday {

							} else {
								fmt.Println("Week Day is Saturday!!")
							}
						} else {
							fmt.Println("Week Day is Friday!!")
						}
					} else {
						fmt.Println("Week Day is Thursaday!!")
					}
				} else {
					fmt.Println("Week Day is Wednsday!!")
				}
			} else {
				fmt.Println("Weekday is Tuesday!!")
			}
		} else {
			fmt.Println("Weekday is Monday!!")
		}
	} else {
		fmt.Println("Weekday is Sunday!!")
	}

}
