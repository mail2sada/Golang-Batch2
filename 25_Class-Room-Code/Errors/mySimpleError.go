package main

import "fmt"

type SimpleError int

func (s SimpleError) Error() string {

	return fmt.Sprintf("Error is %d", s)
}

func Divide(a, b int) (ret float64, err error) {
	if b == 0 {
		var myError SimpleError = 10
		err = &myError
		return
	}

	ret = float64(a) / float64(b)
	return
}

func main() {

	fmt.Println("Demo: Custom Error...")
	ret, err := Divide(10, 0)

	if err != nil {
		var a *SimpleError = err.(*SimpleError)
		fmt.Println(err)
		fmt.Println(*a)
		return
	}
	fmt.Println(ret)
}
