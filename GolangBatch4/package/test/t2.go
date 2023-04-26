package test

import "fmt"

func init() {
	fmt.Println("This init is from t2 from test package")
}

func TestCode() {
	fmt.Println("Beore calling test1")
	test1()
}
