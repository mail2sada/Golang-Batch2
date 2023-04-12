package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: What is today...")

	var t = time.Now()

	fmt.Println("Todays is:", t)

	fmt.Println("Today is ", t.Day(), "Of month")
	fmt.Println("Current month is:", t.Month(), "")
	fmt.Println("Current hour is:", t.Hour())
	var t1 = time.Date(2023, 3, 31, 9, 15, 0, 0, &time.Location{})
	fmt.Println(t1)
	fmt.Println("Difference between:", t1.Sub(t))

	var t2 = t.Add(45 * time.Minute)
	fmt.Println("Time difference is:", t2.Sub(t))

	var format = t.Format("Monday 02-January-2006 03:04:05:.000s .000000ms .000000000ns PM -0700")
	fmt.Println(format)
	//"Friday 31-March-2023 09:44:42:.171s .171051ms .171051000ns AM +0530"

	t5, err := time.Parse("Monday 02-January-2006 03:04:05:.000s .000000ms .000000000ns PM -0700", "Friday 31-March-2023 09:44:42:.171s .171051ms .171051000ns AM +0530")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t5)

	loc, err := time.LoadLocation("Asia/Calcutta")

	if err != nil {

	}

	tx := time.Date(2023, 03, 31, 9, 50, 0, 0, loc)

	fmt.Println(tx)
}
