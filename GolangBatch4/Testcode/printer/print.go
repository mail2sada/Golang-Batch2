package printer

type Printer interface {
	Print()
	Preview()
	PageSetup()
}

func PrintFile(p Printer) {
	p.Print()
}

func PreviewFile(p Printer) {
	p.Preview()
}

func PageSetup(p Printer) {
	p.PageSetup()
}
