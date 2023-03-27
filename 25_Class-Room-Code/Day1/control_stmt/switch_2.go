package Mavenir

import (
	"fmt"
	"time"
)

func main() {

	switch weekDay := time.Now().Weekday(); weekDay {
	case time.Sunday:
	case time.Monday:
	case time.Tuesday:
	case time.Wednesday:
	case time.Thursday:
	case time.Friday:
	case time.Saturday:
	}

	month := "March"
	switch month {
	case "March":
		fmt.Println("March...")

	default:
		fmt.Println("I dont know")
	}

	switch  {
	case time.Now().Day() < 7:
		fmt.Println("First week")
	case time.Now().Day() < 15:
		fmt.Println()
	case time.Month() == 3
		
	}
}
