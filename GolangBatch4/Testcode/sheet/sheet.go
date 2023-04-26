package sheet

import "fmt"

type SpreadSheet struct {
	SheetID   int
	SheetName string
}

func (d SpreadSheet) Print() {
	fmt.Println("In SpreadSheet Print")
}

func (d SpreadSheet) PageSetup() {
	fmt.Println("In SpreadSheet PageSetup")
}

func (d SpreadSheet) Preview() {
	fmt.Println("In SpreadSheet Preview")
}
