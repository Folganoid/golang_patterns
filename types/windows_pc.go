package types

import (
	"../base"
)

type WinPC struct {
	scanner base.Scanner
}

func (pc *WinPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *WinPC) Scan() {
	println("Scan pc Windows system")
	pc.scanner.ScanFile()
}