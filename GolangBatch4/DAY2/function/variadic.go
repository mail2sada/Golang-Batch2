package main

import "fmt"

func Variadic(str ...string) {
	for _, arg := range str {
		fmt.Print(arg)
	}
}

func main() {
	fmt.Println("Demo", "Variadic arguments", 10, 10.111)
	Variadic("Hello", "123", "456", "789")

	ans := Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("\n", ans)
	ans = Add()
	fmt.Println(ans)

	ans = Add(1)
	fmt.Println(ans)

}
func Add(ints ...int) int {
	res := 0
	for _, intVal := range ints {
		res += intVal
	}
	return res
}
