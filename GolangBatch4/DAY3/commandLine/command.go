package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	fmt.Println(args)

	count := flag.Int("count", 0, "represents total count")
	name := flag.String("name", "", "name of arguments")

	flag.Parse()

	// if *count == 0 && *name == "" {
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	str := flag.Args()
	fmt.Println(str)
	fmt.Println(*count)
	fmt.Println(*name)

}
