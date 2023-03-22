package main

import "fmt"

func main() {
	fmt.Println("Demo: Rune...")

	var rune0 rune = '♛'
	var rune1, rune2, rune3, rune4 rune = '♠', '♧', '♡', '♬'

	fmt.Println(rune0, rune1, rune2, rune3, rune4)

	fmt.Printf("%c %c %c %c %c", rune0, rune1, rune2, rune3, rune4)

}
