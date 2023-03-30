package main

import (
	"fmt"
)

type MyCustomError struct {
	errorCode   int
	errorString string
}

func (err *MyCustomError) Error() string {
	res := fmt.Sprintf("ErrorCode:%d and Error String:%s", err.errorCode, err.errorString)
	return res
}

func DivideSlice(x []int, y []int) (res []float64, err error) {

	if x == nil || y == nil {
		customError := new(MyCustomError)
		customError.errorCode = 1
		customError.errorString = "One of the slice is not initialised"
		err = customError
		return
	}

	if len(x) == 0 || len(y) == 0 {
		customError := new(MyCustomError)
		customError.errorCode = 2
		customError.errorString = "One of the slice is empty"
		err = customError
		return
	}

	if len(x) != len(y) {
		customError := new(MyCustomError)
		customError.errorCode = 3
		customError.errorString = "Should have same number elements in both slices"
		err = customError
		return
	}
	res = make([]float64, len(x), len(x))
	for i := 0; i < len(x); i++ {
		if y[i] == 0 {
			customError := new(MyCustomError)
			customError.errorCode = 4
			customError.errorString = fmt.Sprintf("Y at index %d is zero", i)
			err = customError
			return
		}
		res[i] = float64(x[i]) / float64(y[i])
	}

	return
}

func main() {
	fmt.Println("Demo: Custom errors")

	sliceX := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slicey := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//var slicey []int

	res, err := DivideSlice(sliceX, slicey)

	if err != nil {
		var customError *MyCustomError = err.(*MyCustomError)
		switch customError.errorCode {
		case 0:
			fmt.Println("Undefined")
		case 1:
			fmt.Println("Error 1")
		case 2:
			fmt.Println("Error 2")
		case 3:
			fmt.Println("Error 3")
		default:
			fmt.Println("Undefined")
		}
		fmt.Println(customError.errorString)

	} else {
		fmt.Println(res)
	}
}
