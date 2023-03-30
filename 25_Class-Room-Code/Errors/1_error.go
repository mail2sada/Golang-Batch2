package main

import (
	"errors"
	"fmt"
)

func Divde(x, y int) (err error, res int) {
	if y == 0 {
		err = errors.New("Divide by zero...")
		return
	}
	res = x / y
	return
}

func MyDivide(x, y int) (res float64, err error) {
	if y == 0 {
		err = fmt.Errorf("Divide by zero %d/%d expecting non zero value for y", x, y)
		return
	}
	res = float64(x) / float64(y)
	return
}

func main() {
	fmt.Println("Demo: Errors")
	err, res := Divde(10, 20)

	fmt.Println(res, err)

	err, res = Divde(100, 0)

	fmt.Println(res, err)

	result, err := MyDivide(1000, 0)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
