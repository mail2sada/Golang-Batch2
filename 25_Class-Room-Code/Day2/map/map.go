package main

import "fmt"

func main() {

	var url map[string]string = map[string]string{}

	var emp map[string]int = make(map[string]int)

	url["wrg"] = "https://wrg.mavenir.com"
	url["m-store"] = "https://m-store.mavenir.com"

	str, non := url["wrg"]

	if non {
		fmt.Println("Key not found")
	}

	fmt.Println(str)
	fmt.Println(url["wrg"])

	fmt.Println(url["m-store"])

	fmt.Println(url)
}
