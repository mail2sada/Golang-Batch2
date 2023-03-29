package main

import "fmt"

func Divide(x, y int) (bool, float64) {

	err := false
	ans := 0.0
	if y == 0 {
		err = true
		return err, ans
	}
	ans = float64(x) / float64(y)
	return err, ans
}

func main() {
	fmt.Println("Demo func with multiple return")

	fmt.Println(Divide(20, 10))

	fmt.Println(Divide(200, 10))

	fmt.Println(Divide(20, 0))

}
