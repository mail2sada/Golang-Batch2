package main

import "fmt"

func main() {

	fmt.Println("Demo: Maps")

	var mmap map[string]int = make(map[string]int)

	mmap["A"] = 10

	mmap["B"] = 20

	mmap["C"] = 30

	for key, val := range mmap {
		fmt.Println("Key is ", key)
		fmt.Println("Value is", val)
	}

	// Reset map
	// for key, _ := range mmap {
	// 	delete(mmap, key)
	// }

	mmap = make(map[string]int)

	fmt.Println(mmap)

}
