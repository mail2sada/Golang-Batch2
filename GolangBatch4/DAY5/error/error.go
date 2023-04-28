package main

import (
	"fmt"
)

func Divide(a, b int) (res int, err error) {

	if b == 0 {
		//err = errors.New("Divide by zero")

		err = fmt.Errorf("divide by zero error received")
		return
	}

	res = a / b

	return
}

func main() {

	a, _ := Divide(20, 10)

	fmt.Println(a)

	b, _ := Divide(50, 5)

	fmt.Println(b)

	c, e := Divide(10, 0)

	if e != nil {
		fmt.Println("Received error while performing divide", e)
	}

	fmt.Println(c)

}
