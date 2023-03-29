package main

import (
	"fmt"
	"sort"
)

func main() {
	var emp map[string]string = map[string]string{"Anil": "NGN",
		"Vijay":  "MDE",
		"Manoj":  "CGFM",
		"Deepak": "Radio"}

	// Read all keys in a slice
	var keySlice []string = []string{}

	for key, _ := range emp {
		keySlice = append(keySlice, key)
	}
	fmt.Println(keySlice)
	sort.Strings(keySlice)
	fmt.Println(keySlice)

	for _, key := range keySlice {
		fmt.Println("Key is ", key, "Value is", emp[key])
	}
	// sort slice
	// print using slice

}
