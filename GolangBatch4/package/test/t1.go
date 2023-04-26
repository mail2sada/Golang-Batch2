package test

import "fmt"

func init() {
	fmt.Println("This init is from t1 of test package")
}

func test1() {
	fmt.Println("This is test function")
}
