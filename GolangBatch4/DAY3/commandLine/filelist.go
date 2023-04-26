package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Demo: File list....")
	dir := flag.String("dir", "./", "root directory to list files")
	flag.Parse()
	list, err := os.ReadDir(*dir)

	if err != nil {
		fmt.Println("Received error while reading directory", err)
		os.Exit(1)
	}
	for _, val := range list {
		fmt.Println(val.Name())
	}
}
