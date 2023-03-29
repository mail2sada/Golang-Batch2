package main

import "fmt"

func Add(integers ...int) (ans int) {

	for _, val := range integers {
		ans += val
	}
	return
}

func Divide(x, y int) (err bool, res float64) {
	if y == 0 {
		err = true
		return
	}
	res = float64(x) / float64(y)
	return
}

func main() {

	fmt.Println("Demo: Named return..")

	fmt.Println(Add(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))

	fmt.Println(Divide(10, 20))
	fmt.Println(Divide(100, 20))

	fmt.Println(Divide(1000, 20))

	fmt.Println(Divide(100000, 0))

}
