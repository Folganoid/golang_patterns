package types

import (
	"../base"
)

type MacPC struct {
	scanner base.Scanner
}

func (pc *MacPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *MacPC) Scan() {
	println("Scan pc Mac system")
	pc.scanner.ScanFile()
}