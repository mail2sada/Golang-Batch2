package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Date:::")
	t := time.Now()
	fmt.Println("Current time is :", t)
	tAfter180 := t.AddDate(0, 0, 180)
	duration := tAfter180.Sub(t)
	fmt.Println(duration)
	newTime := time.Date(2100, 12, 31, 23, 59, 59, 0, &time.Location{})
	timeString := newTime.Format("Mon 02/January/06 03:04:05 .000ms .000000us .000000000ns pm MST -0700")

	//"Fri 31/March/23 10:30:30 .689ms .689051us .689051000ns am IST +0530"
	fmt.Println(timeString)

	ts, _ := time.Parse("Mon 02/January/06 03:04:05 .000ms .000000us .000000000ns pm MST -0700", "Fri 31/March/23 10:30:30 .689ms .689051us .689051000ns am IST +0530")

	fmt.Println(ts)
}
