package main

import (
	"fmt"
	"interface/doc"
	"interface/image"
	"interface/printer"
	"interface/sheet"
)

func main() {

	fmt.Println("Printer Interface")

	var p1, p2, p3 printer.Printer

	p1 = doc.Document{}
	p2 = sheet.SpreadSheet{}
	p3 = image.Image{}
	p1.Print()
	p2.PageSetup()
	p3.Preview()
}
