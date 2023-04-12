package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo time conversion")

	currenttime := time.Now()
	// go time epoch
	unixTime := currenttime.Unix()

	fmt.Println(unixTime)
	//epoch to go time
	tm := time.Unix(unixTime, 0)

	fmt.Println(tm)
}
