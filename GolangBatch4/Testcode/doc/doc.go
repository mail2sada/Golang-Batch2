package doc

import "fmt"

type Document struct {
	DocId   int
	DocName string
}

func (d Document) Print() {
	fmt.Println("In Document Print")
}

func (d Document) PageSetup() {
	fmt.Println("In Document PageSetup")
}

func (d Document) Preview() {
	fmt.Println("In Document Preview")
}
