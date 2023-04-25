package main

import "fmt"

func main() {
	fmt.Println("Demo: Maps")

	var firstMap map[string]string = map[string]string{"A": "Alpha", "B": "Beta"}

	val := firstMap["A"]
	fmt.Println(val)
	firstMap["D"] = "Delta"
	firstMap["G"] = "Gama"
	firstMap["N"] = "Nano"
	firstMap["O"] = "OCta"
	fmt.Println(firstMap)

	val1, aval := firstMap["G"]
	if aval {
		fmt.Println(val1)
	} else {
		fmt.Println("Unknown key")
	}

	for key, val := range firstMap {
		fmt.Println("Key:", key, " Value:", val)
	}

	delete(firstMap, "G")

	fmt.Println(firstMap)

}
