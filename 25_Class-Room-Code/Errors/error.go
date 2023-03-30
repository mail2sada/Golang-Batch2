package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Demo: Errors")
	file, err := os.OpenFile("File1.txt", os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println("Error is %w", err)
	}

	file.Write([]byte("ABC"))
}
