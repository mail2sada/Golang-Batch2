package main

import "fmt"

func main() {
	fmt.Println("Demo loop")
label1:
	var i = 0

	for i = 0; i < 10; i++ {
		fmt.Println("Value of i", i)
	}
	for i < 20 {
		fmt.Println("i = ", i)
		i++
	}
	for {
		//fmt.Println("i", i)
		i++
		if i < 100 {
			goto label1
		}
		fmt.Println("Value of i is ", i)

		if i == 200 {
			break
		}

	}
}
