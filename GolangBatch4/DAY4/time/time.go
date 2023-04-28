package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t)

	//t1 := time.Date(2023, time.April, 27, 15, 0, 0, 0, &time.Location{})

	// t2 := t.Add(45 * time.Minute)

	// dur := t.Sub(t2)

	// fmt.Println(dur)
	// fmt.Println(t2)

	seconds := t.Unix()

	fmt.Println(seconds)

	t = time.Unix(1682587860, 0)
	fmt.Println(t)

	str := t.Format("2/Jan/2006  15:4:05")

	fmt.Println(str)

	t, _ = time.Parse("2/Jan/2006  15:4:05", "27/Apr/2023  15:1:00")
	fmt.Println(t)
}
