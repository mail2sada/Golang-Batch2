package main

import "fmt"

type WeekDay uint8

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednsday
	Thursaday
	Friday
	Saturday
)

func main() {

	switch dayOfWeek := Thursaday; dayOfWeek {
	case Monday, Tuesday, Wednsday, Thursaday, Friday:
		fmt.Println("Working Day...")

	case Saturday, Sunday:
		fmt.Println("Weekends funtime...")
	}
}
