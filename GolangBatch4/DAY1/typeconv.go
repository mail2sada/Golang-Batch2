package main

import (
	"fmt"
	"reflect"
)

func main() {
	var m1, m2, m3 = 100, 99, 99

	//var avg float64

	avg := float64(m1+m2+m3) / 3
	fmt.Println("Average marks", avg)
	fmt.Println("Type of avg ", reflect.TypeOf(avg))
}
