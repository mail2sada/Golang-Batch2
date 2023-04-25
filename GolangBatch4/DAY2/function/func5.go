package main

import (
	"errors"
	"fmt"
)

func Math(a, b int) (sum, diff, prod, div int, err error) {
	sum = a + b
	diff = a - b
	prod = a * b
	if b == 0 {
		err = errors.New("divide by zero error")
		return
	}
	div = a / b
	return
}

func main() {

	a, b, c, d, e := Math(10, 0)

	fmt.Println(a, b, c, d, e)

}
