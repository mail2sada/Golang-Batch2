package main

import (
	"fmt"
)

func main() {

	percentageMarks := 90

	switch {
	case percentageMarks < 50:
		fmt.Println("Better luck next time")
	case percentageMarks < 60:
		fmt.Println("You just made it, it could have been still better")
	case percentageMarks < 70:
		fmt.Println("You did Ok...")
	case percentageMarks <= 80:
		fmt.Println("You have done well")
	case percentageMarks > 80:
		fmt.Println("Congratulations, you have performed very well....")
	}
}
