package main

import "fmt"

func main() {
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	winComputer := Windows{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	fmt.Println()
}
