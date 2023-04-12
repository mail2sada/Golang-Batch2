package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Date and time..")
	var t time.Time = time.Now()

	fmt.Println("Current time is:", t)
	fmt.Println("Todays date is:", t.Day())
	fmt.Println("Current month is", t.Month())
	fmt.Println("Todays Weekday:", t.Weekday())
	fmt.Println("Current Hour is:", t.Hour())
	fmt.Println("Current Minute is:", t.Minute())
	fmt.Println("Current Second is:", t.Second())

	var t1 time.Time = time.Date(2023, 3, 31, 10, 10, 0, 0, &time.Location{})
	fmt.Println("Date constructed:", t1)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	var t2 time.Time = time.Date(2023, 3, 31, 10, 10, 0, 0, loc)
	fmt.Println("Date constructed UTC:", t2)

	// Adding duration to time.

	var tAfter = t.Add(45 * 60 * 1000 * 1000 * 1000)
	fmt.Println(tAfter)

	//var tAf = t.Add((45 * time.Minute))

	var tAddDate = t.AddDate(0, 0, 180)

	fmt.Println(tAddDate)

}
