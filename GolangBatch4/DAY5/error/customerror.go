package main

import (
	"fmt"
)

type CustomError struct {
	ErrorCode    int
	ErrorMessage string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("*#ErrorCode:[%d] and ErrorMessage[%s]", c.ErrorCode, c.ErrorMessage)
}

func Divide(a, b int) (res int, err error) {

	if b == 0 {
		customError := &CustomError{ErrorCode: 100, ErrorMessage: "Divide by zero"}
		err = customError
		return
	}
	res = a / b
	return
}
func main() {
	fmt.Println("Custom error")

	a, err := Divide(100, 0)
	if err != nil {
		fmt.Println("got error", err)

		cerr := err.(*CustomError)
		fmt.Println(cerr.ErrorCode)
		fmt.Println(cerr.ErrorMessage)
	} else {
		fmt.Println("Result is", a)
	}

}

func CheckThis() {
	var a interface{}

	a = 10
	fmt.Println(a)

	a = "Hello"

	fmt.Println(a)

	switch a.(type) {
	case int:
		intval := a.(int)
		fmt.Println(intval)
	case string:
		strval := a.(int)
		fmt.Println(strval)
	}
}
