package types

import (
	"../base"
)

type LinuxPC struct {
	scanner base.Scanner
}

func (pc *LinuxPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *LinuxPC) Scan() {
	println("Scan pc Linux system")
	pc.scanner.ScanFile()
}