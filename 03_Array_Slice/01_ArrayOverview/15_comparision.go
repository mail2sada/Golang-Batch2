package main

import "fmt"

func main() {
	fmt.Println("Demo: comparision of arrays")

	var stringArray = [5]string{"Hello", "how ", "are", "you", "?"}

	fmt.Println("Contents of stringArray:", stringArray)

	fmt.Println("Copying array...")
	copyArray := stringArray

	if copyArray == stringArray {
		fmt.Println("Both arrays are equal..")
	}

	if copyArray != stringArray {
		fmt.Println("Both arrays are unequal..")
	}

	fmt.Println("Changing contents of array..")

	stringArray[4] = "Wow"

	if copyArray == stringArray {
		fmt.Println("Both arrays are equal..")
	}

	if copyArray != stringArray {
		fmt.Println("Both arrays are unequal..")
	}

}
