package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Demo: timer...")

	c := time.After(2 * time.Second)

	val := <-c

	fmt.Println(val, "Timer is expired")

}
