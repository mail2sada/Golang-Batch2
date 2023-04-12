package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Demo Os")

	//os.Mkdir("Test", 0644)
	os.Chdir("Test")

	file, err := os.OpenFile("./test.txt", os.O_CREATE, 0777)

	if err != nil {
		fmt.Println("Failed to create...", err)
	}
	defer file.Close()
	fmt.Fprintln(file, "This is our text file...")
}
