package image

import "fmt"

type Image struct {
	ImageId int
	Name    string
}

func (d Image) Print() {
	fmt.Println("In Image Print")
}

func (d Image) PageSetup() {
	fmt.Println("In Image PageSetup")
}

func (d Image) Preview() {
	fmt.Println("In Image Preview")
}
