// Write a command line tools, that accepts filename with flag -flag , count with -count integert, repeat boolean
package main

import (
	"flag"
	"fmt"
)

type MyList []string

func (lst *MyList) String() string {
	return fmt.Sprintln(*lst)
}

func (lst *MyList) Set(s string) error {
	*lst = append(*lst, s)
	return nil
}

var (
	filename = flag.String("file", "", "file-name")
	count    = flag.Int("count", 2, "Count for the iteration")
	repeat   = flag.Bool("repeat", false, "Boolean value for repeat")
)

func main() {
	var list MyList
	flag.Var(&list, "list", "list of files")
	flag.Parse()
	fmt.Println("Demo: Commmand line arguments")

	fmt.Println("-file is passed", *filename)
	fmt.Println("-count is ", *count)

	fmt.Println(list)

	fmt.Println(flag.Args())

}
